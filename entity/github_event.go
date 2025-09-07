package entity

type GithubEvent struct {
	Type      string `json:"type"`
	IsPublic  bool   `json:"public"`
	CreatedAt string `json:"created_at"`
}
