package youtubeComments

const defaultUserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36"

const baseUrl = "https://www.youtube.com"
const videoKey = "{videoKey}"
const videoUrlTemplate = baseUrl + "/watch?v=" + videoKey
const videoCommentsURL = baseUrl + "/comment_service_ajax"

const getCommentsAction = "action_get_comments"
const getRepliesAction = "action_get_comment_replies"
