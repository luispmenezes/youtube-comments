package youtubeComments

import (
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
	"youtube-comments/pkg/youtubeComments/external"
	"youtube-comments/pkg/youtubeComments/model"
)

func findInHtmlString(htmlStr string, key string, offset int, delimiter string) string {
	log.WithFields(log.Fields{"key": key}).Debug("findInHtmlString")
	startIdx := strings.Index(htmlStr, key) + len(key) + offset
	endIdx := startIdx + strings.Index(htmlStr[startIdx:], delimiter)
	return htmlStr[startIdx:endIdx]
}

func unescapeUnicode(source string) string {
	unescapedStr, err := strconv.Unquote(`"` + source + `"`)
	if err != nil {
		log.Fatal(err)
	}
	return unescapedStr
}

func findInMap(key string, sourceMap map[string]interface{}) interface{} {
	for k, v := range sourceMap {
		if k == key {
			return v
		}
		switch v.(type) {
		case map[string]interface{}:
			if foundValue := findInMap(key, v.(map[string]interface{})); foundValue != nil {
				return foundValue
			}
			break
		case []interface{}:
			for _, m := range v.([]interface{}) {
				switch m.(type) {
				case map[string]interface{}:
					if foundValue := findInMap(key, m.(map[string]interface{})); foundValue != nil {
						return foundValue
					}
					break
				}
			}
			break
		}
	}
	return nil
}

func buildComment(commentRenderer external.CommentRenderer) model.Comment {
	author := model.Author{
		Name:            commentRenderer.AuthorText.SimpleText,
		ChannelEndpoint: commentRenderer.AuthorEndpoint.BrowseEndpoint.CanonicalBaseUrl,
		Thumbnails:      []model.Thumbnail{},
		IsChannelOwner:  commentRenderer.AuthorIsChannelOwner,
	}

	for _, thumbnail := range commentRenderer.AuthorThumbnail.Thumbnails {
		author.Thumbnails = append(author.Thumbnails, model.Thumbnail{
			URL:    thumbnail.URL,
			Width:  thumbnail.Width,
			Height: thumbnail.Height,
		})
	}

	commentText := ""
	for _, run := range commentRenderer.ContentText.Runs {
		commentText += run.Text
	}

	timeString := ""
	var age time.Duration

	if len(commentRenderer.PublishedTimeText.Runs) > 0 {
		timeString = commentRenderer.PublishedTimeText.Runs[0].Text
		timeStringTokens := strings.Split(timeString, " ")
		if len(timeStringTokens) >= 2 {
			amount, err := strconv.Atoi(timeStringTokens[0])
			if err != nil {
				log.Fatal(err)
			}

			if strings.HasPrefix(timeStringTokens[1], "year") {
				age = time.Duration(amount*365*24) * time.Hour
			} else if strings.HasPrefix(timeStringTokens[1], "month") {
				age = time.Duration(amount*30*24) * time.Hour
			} else if strings.HasPrefix(timeStringTokens[1], "week") {
				age = time.Duration(amount*7*24) * time.Hour
			} else if strings.HasPrefix(timeStringTokens[1], "day") {
				age = time.Duration(amount*24) * time.Hour
			} else if strings.HasPrefix(timeStringTokens[1], "hour") {
				age = time.Duration(amount) * time.Hour
			}
		}
	}

	comment := model.Comment{
		Id:          commentRenderer.CommentId,
		CommentText: commentText,
		LikeCount:   commentRenderer.LikeCount,
		IsHearted:   commentRenderer.ActionButtons.CommentActionButtonsRenderer.CreatorHeart.CreatorHeartRenderer.IsHearted,
		IsEdited:    strings.Contains(timeString, " (edited)"),
		Age:         age,
		Author:      author,
	}
	return comment
}
