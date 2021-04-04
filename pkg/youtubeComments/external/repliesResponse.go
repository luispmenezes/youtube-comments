package external

type RepliesResponse []struct {
	Page      string      `json:"page,omitempty"`
	XSRFToken string      `json:"xsrf_token,omitempty"`
	Endpoint  interface{} `json:"endpoint,omitempty"`
	Response  Response    `json:"response,omitempty"`
	Timing    Timing      `json:"timing,omitempty"`
}