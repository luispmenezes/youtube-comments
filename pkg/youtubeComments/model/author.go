package model

type Author struct {
	Name            string      `json:"name"`
	ChannelEndpoint string      `json:"channelEndpoint"`
	Thumbnails      []Thumbnail `json:"thumbnail"`
	IsChannelOwner  bool        `json:"isChannelOwner"`
}

type Thumbnail struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
