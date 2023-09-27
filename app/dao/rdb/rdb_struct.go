package rdb

type AuthorizeCache struct {
	FromProvider string `json:"from_provider"`
	State        string `json:"state"`
	ClientId     string `json:"client_id"`
	RedirectUri  string `json:"redirect_uri"`
}
