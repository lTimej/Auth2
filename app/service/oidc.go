package service

import (
	"auth2/app/dao/auth"
	"auth2/app/dao/rdb"
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type OIDCService struct {
}

func NewOIDCService() *OIDCService {
	return &OIDCService{}
}

type AuthorizeRequest struct {
	FromProvider string `json:"from_provider"`
	State        string `json:"state"`
	ClientId     string `json:"client_id"`
	RedirectUri  string `json:"redirect_uri"`
	ResponseType string `json:"response_type"`
}

type GrantAndRedirectRequest struct {
	Code  string `json:"code"`
	State string `json:"state"`
}

func (oidc *OIDCService) Authorize(data AuthorizeRequest) {
	client_id := data.ClientId
	pa := auth.GetProviderApplicationByClienId(client_id)
	if pa == nil {
		panic("client_id error")
	}
	ctx := context.Background()
	err := rdb.RDB.Exists(ctx, data.State).Err()
	if err != redis.Nil {
		panic("state exist")
	}
	ac := rdb.AuthorizeCache{
		FromProvider: data.FromProvider,
	}
	rdb.RDB.Set(context.Background(), data.State, ac, time.Second*60)
	oidc := auth.GetODICProviderByProvideCode(data.FromProvider)
	fmt.Println(oidc)
}

func (oidc *OIDCService) GrantAndRedirect(data GrantAndRedirectRequest) {
	state := data.State
	_, err := rdb.RDB.Get(context.Background(), state).Result()
	if err == redis.Nil {
		panic("state session expire")
	} else if err != nil {
		panic(err)
	}

}
