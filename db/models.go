package db

type (
	Request struct {
		id      int               `json:"id,omitempty"`
		Url     string            `json:"url,omitempty"`
		Headers map[string]string `json:"headers,omitempty"`
		Body    string            `json:"body,omitempty"`
	}

	Profile struct {
		Email    string `json:"email,omitempty"`
		Name     string `json:"name,omitempty"`
		Password string `json:"password,omitempty"`
	}
)

func (p *Profile) sync() (*string, *string, *string) {
	return &p.Email, &p.Name, &p.Password
}

func (r *Request) sync() (*int, *string, *map[string]string, *string) {
	return &r.id, &r.Url, &r.Headers, &r.Body
}
