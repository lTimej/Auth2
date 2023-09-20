package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(100)"`
	Password string `json:"-" gorm:"type:varchar(100)"`
}

func (u *User) TableName() string {
	return "user"
}

type Client struct {
	gorm.Model
	ClientID     string `json:"client_id" gorm:"type:varchar(100)"`
	ClientSecret string `json:"client_secret" gorm:"type:varchar(255)"`
	RedirectURI  string `json:"redirect_uri" gorm:"type:varchar(500)"`
}

func (c *Client) TableName() string {
	return "client"
}

type Token struct {
	ID           uint      `gorm:"primary_key"`
	UserID       int64     `json:"-"`
	ClientID     string    `json:"client_id" gorm:"type:varchar(100)"`
	ClientSecret string    `json:"client_secret" gorm:"type:varchar(255)"`
	TokenType    string    `json:"token_type" gorm:"type:varchar(32)"`
	ExpiresIn    int32     `json:"expires_in"`
	RefreshToken string    `json:"refresh_token" gorm:"type:varchar(100)"`
	CreateTime   time.Time `json:"create_time"`
	UpdateTime   time.Time `json:"update_time"`
}

func (t *Token) TableName() string {
	return "token"
}

type OpenIDConfig struct {
	gorm.Model
	SchedProviderCode string `json:"sched_provider_code" gorm:"type:varchar(64);unique"  binding:"required"`
	Issuer            string `json:"issuer" gorm:"type:varchar(100)"  binding:"required"`
	ClientID          string `json:"client_id" gorm:"type:varchar(100)"  binding:"required"`
	ClientSecret      string `json:"client_secret" gorm:"type:varchar(255)"  binding:"required"`
	Creator           string `json:"creator" gorm:"type:varchar(150)"`
	Updator           string `json:"updator" gorm:"type:varchar(150)"`
}

func (ocid *OpenIDConfig) TableName() string {
	return "openidconfig"
}

type ProviderApplication struct {
	gorm.Model
	ClientID               string `json:"client_id" gorm:"type:varchar(100)"`
	ClientSecret           string `json:"client_secret" gorm:"type:varchar(255)"`
	RedirectUri            string `json:"redirect_uri" gorm:"type:varchar(500)"`
	ClientType             string `json:"client_type" gorm:"type:varchar(32)"`
	PostLogoutRedirectUri  string `json:"post_logout_redirect_uri" gorm:"type:varchar(500)"`
	AuthorizationGrantType string `json:"authorization_grant_type" gorm:"type:varchar(32)"`
	Name                   string `json:"name" gorm:"type:varchar(255)"`
	SkipAuthorization      bool   `json:"skip_authorization"`
	Algorithm              string `json:"algorithm" gorm:"type:varchar(5)"`
	ProviderCode           string `json:"provider_code" gorm:"type:varchar(64)"`
	UserId                 uint   `json:"user_id" gorm:"type:int"`
	User                   *User  `gorm:"ForeignKey:UserId;AssociationForeignKey:ID;constraint:OnDelete:CASCADE"`
}

func (pa *ProviderApplication) TableName() string {
	return "providerapplication"
}
