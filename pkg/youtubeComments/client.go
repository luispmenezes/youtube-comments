package youtubeComments

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
	"youtube-comments/pkg/youtubeComments/external"
	"youtube-comments/pkg/youtubeComments/model"
)

type YoutubeCommentsClient struct {
	jar             *cookiejar.Jar
	httpClient      *http.Client
	commentCallback CommentCallback
}

func NewYoutubeCommentsClient() *YoutubeCommentsClient {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal("Bateu")
	}

	var cookies []*http.Cookie
	cookie := &http.Cookie{
		Name:   "CONSENT",
		Value:  "YES+cb",
		Path:   "/",
		Domain: ".youtube.com",
		Secure: true,
	}
	cookies = append(cookies, cookie)
	u, _ := url.Parse(baseUrl)
	jar.SetCookies(u, cookies)

	return &YoutubeCommentsClient{
		jar: jar,
		httpClient: &http.Client{
			Jar: jar,
		}}
}

func (yc *YoutubeCommentsClient) GetComments(videoId string) ([]model.Comment, error) {
	log.Debug("GetComments ", videoId)
	videoUrl := strings.Replace(videoUrlTemplate, videoKey, videoId, 1)
	statusCode, body, err := yc.sendRequest(http.MethodGet, videoUrl, nil, nil)

	if err != nil {
		return nil, errors.Wrap(err, "failed to get watch page")
	}

	if statusCode == http.StatusOK {
		sessionToken := unescapeUnicode(findInHtmlString(body, "XSRF_TOKEN", 3, "\""))

		dataStr := findInHtmlString(body, "var ytInitialData = ", 0, "};") + "}"
		var dataMap map[string]interface{}
		err = json.Unmarshal([]byte(dataStr), &dataMap)
		if err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal datamap from html")
		}

		continuationData := findInMap("nextContinuationData", dataMap)

		if continuationData == nil {
			return nil, errors.New("could not find nextContinuationData in data map")
		}

		continuationDataMap := continuationData.(map[string]interface{})
		if continuationDataMap != nil {
			clickTrackingParams := continuationDataMap["clickTrackingParams"].(string)
			continuation := continuationDataMap["continuation"].(string)
			action := getCommentsAction

			return yc.doCommentsRequest(sessionToken, clickTrackingParams, continuation, action)
		} else {
			log.Debug("No comments")
			return []model.Comment{}, nil
		}
	} else {
		return nil, fmt.Errorf("received bad status code from watch page: %d", statusCode)
	}
}

func (yc *YoutubeCommentsClient) RegisterCommentCallback(callback CommentCallback) {
	yc.commentCallback = callback
}

func (yc *YoutubeCommentsClient) doCommentsRequest(sessionToken, clickTracking, continuation, action string) ([]model.Comment, error) {
	log.WithFields(log.Fields{"sessionToken": sessionToken, "clickTracking": clickTracking,
		"continuation": continuation, "action": action}).
		Debug("doCommentsRequest")

	videoUrl, headers, data, err := yc.prepareCommentsRequest(sessionToken, clickTracking, continuation, action)
	if err != nil {
		return nil, err
	}

	statusCode, body, err := yc.sendRequest(http.MethodPost, videoUrl.String(), strings.NewReader(data.Encode()), headers)

	if err != nil {
		return nil, errors.Wrap(err, "failed to send comments ajax")
	}

	var commentList []model.Comment

	if statusCode == http.StatusOK {
		var response external.Response
		if action == getRepliesAction {
			var repliesResponse external.RepliesResponse
			if err := json.Unmarshal([]byte(body), &repliesResponse); err != nil {
				log.Debug(body)
				return nil, errors.Wrap(err, "failed to parse replies ajax response")
			}
			response = repliesResponse[1].Response
		} else {
			var commentResponse external.CommentResponse
			if err := json.Unmarshal([]byte(body), &commentResponse); err != nil {
				log.Debug(body)
				return nil, errors.Wrap(err, "failed to parse comments ajax response")
			}
			response = commentResponse.Response
		}

		for _, thread := range response.ContinuationContents.ItemSectionContinuation.Contents {
			comment := buildComment(thread.CommentThreadRenderer.Comment.CommentRenderer)
			for _, c := range thread.CommentThreadRenderer.Replies.CommentRepliesRenderer.Continuations {
				nextContinuation := c.NextContinuationData
				replies, err := yc.doCommentsRequest(sessionToken, nextContinuation.ClickTrackingParams,
					nextContinuation.Continuation, getRepliesAction)

				if err != nil {
					return nil, errors.Wrap(err, "error getting replies")
				}

				comment.Replies = append(comment.Replies, replies...)
			}
			commentList = append(commentList, comment)
		}

		if len(commentList) > 0 && yc.commentCallback != nil {
			yc.commentCallback(commentList)
		}

		for _, c := range response.ContinuationContents.ItemSectionContinuation.Continuations {
			nextContinuation := c.NextContinuationData
			contComments, err := yc.doCommentsRequest(sessionToken, nextContinuation.ClickTrackingParams,
				nextContinuation.Continuation, getCommentsAction)

			if err != nil {
				return nil, errors.Wrap(err, "error getting continuation comments")
			}

			commentList = append(commentList, contComments...)
		}

	} else {
		return nil, fmt.Errorf("bad status code for comments response %d", statusCode)
	}

	return commentList, nil
}

func (yc *YoutubeCommentsClient) prepareCommentsRequest(sessionToken, clickTracking, continuation, action string) (*url.URL, map[string]string, url.Values, error) {
	commentUrl, err := url.Parse(videoCommentsURL)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed to parse comments ajax url")
	}

	query := commentUrl.Query()
	query.Set(action, "1")
	query.Set("pbj", "1")
	query.Set("ctoken", continuation)
	query.Set("continuation", continuation)
	query.Set("itct", clickTracking)
	commentUrl.RawQuery = query.Encode()

	headers := map[string]string{}
	headers["x-youtube-client-name"] = "1"
	headers["x-youtube-client-version"] = "2.20210330.08.00"
	headers["accept-language"] = "en-US,en;q=0.9"

	data := url.Values{}
	data.Set("session_token", sessionToken)

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Content-Length"] = strconv.Itoa(len(data.Encode()))

	return commentUrl, headers, data, nil
}

func (yc *YoutubeCommentsClient) sendRequest(verb, url string, body io.Reader, headers map[string]string) (int, string, error) {
	log.WithFields(log.Fields{"verb": verb, "url": url}).Debug("send request")
	req, err := http.NewRequest(verb, url, body)
	if err != nil {
		return 0, "", err
	}

	req.Header.Set("user-agent", defaultUserAgent)

	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}

	resp, err := yc.httpClient.Do(req)
	if err != nil {
		return 0, "", err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, "", err
	}
	bodyString := string(bodyBytes)

	return resp.StatusCode, bodyString, nil
}
