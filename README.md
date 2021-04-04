# Youtube Comments

(Under-Construction)

Get comments from Youtube videos without using the API.

## Usage

### CLI
```
./youtube-comments -videoId dQw4w9WgXcQ -outputPath my-file.json
```

### Lib
```
go get github.com/luispmenezes/youtube-comments
```

```go
package main

import (
	"fmt"
	"github.com/luispmenezes/youtube-comments/pkg/youtubeComments"
)

func main() {
	client := youtubeComments.NewYoutubeCommentsClient()
	comments, err := client.GetComments("dQw4w9WgXcQ")
	if err != nil {
		panic(err)
	}
	fmt.Println(comments)
}
```
With a new comment continuation callback
```go
package main

import (
	"fmt"
	"github.com/luispmenezes/youtube-comments/pkg/youtubeComments"
	"github.com/luispmenezes/youtube-comments/pkg/youtubeComments/model"
)

func main() {
	client := youtubeComments.NewYoutubeCommentsClient()
	client.RegisterCommentCallback(onNewComments)
	client.GetComments("1hgzRK-9J_A")
}

func onNewComments(comments []model.Comment)  {
	fmt.Println(comments)
}
```