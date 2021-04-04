# Youtube Comments

(Under-Construction)

Get comments from Youtube videos without using the API.

## Usage

### CLI
```
./youtube-comments -videoId dQw4w9WgXcQ -outputPath my-file.json
```

### Lib

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