package service

import (
	"auth2/app/dao/auth"
	"auth2/app/dao/rdb"
	"auth2/constants"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	OIDC_KEY string = "state:"
)

type OIDCService struct {
}

func NewOIDCService() *OIDCService {
	return &OIDCService{}
}

type AuthorizeRequest struct {
	FromProvider string `form:"from_provider" json:"from_provider" binding:"required"`
	State        string `form:"state" json:"state" binding:"required"`
	ClientId     string `form:"client_id" json:"client_id" binding:"required"`
	RedirectUri  string `form:"redirect_uri" json:"redirect_uri" binding:"required"`
	ResponseType string `form:"response_type" json:"response_type" binding:"required"`
}

type GrantAndRedirectRequest struct {
	Code  string `form:"code" json:"code" binding:"required"`
	State string `form:"state" json:"state" binding:"required"`
}

func (oidc *OIDCService) Authorize(data AuthorizeRequest) string {
	key := fmt.Sprintf("%s%s", OIDC_KEY, data.State)
	client_id := data.ClientId
	pa := auth.GetProviderApplicationByClienId(client_id)
	ctx := context.Background()
	err := rdb.RDB.Exists(ctx, key).Err()
	if err != nil && err != redis.Nil {
		panic("state exist")
	}
	ac := rdb.AuthorizeCache{
		FromProvider: data.FromProvider,
		State:        data.State,
		RedirectUri:  pa.RedirectUri,
		ClientId:     pa.ClientID,
	}
	ac_byte, _ := json.Marshal(ac)
	if err := rdb.RDB.Set(context.Background(), key, ac_byte, time.Second*60).Err(); err != nil {
		panic(err)
	}
	oidc_obj := auth.GetODICProviderByProvideCode(data.FromProvider)
	params := url.Values{
		"response_type": []string{"code"},
		"client_id":     []string{oidc_obj.ClientID},
		"redirect_uri":  []string{constants.REDIRECT_URL},
		"state":         []string{data.State},
	}
	url := fmt.Sprintf("%s%s?%s", oidc_obj.Issuer, "/authorize", params.Encode())
	return url
}

func (oidc *OIDCService) GrantAndRedirect(data GrantAndRedirectRequest) {
	key := fmt.Sprintf("%s%s", OIDC_KEY, data.State)
	oidc_byte, err := rdb.RDB.Get(context.Background(), key).Bytes()
	if err == redis.Nil {
		panic("state session expire")
	} else if err != nil {
		panic(err)
	}
	var ac rdb.AuthorizeCache
	json.Unmarshal(oidc_byte, &ac)
	oidc_obj := auth.GetODICProviderByProvideCode(ac.FromProvider)
	fmt.Println(oidc_obj)
}
