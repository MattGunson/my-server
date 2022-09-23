package db

type Request struct {
	id      string            `json:"id,omitempty"`
	Url     string            `json:"url,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
	Body    string            `json:"body,omitempty"`
}
