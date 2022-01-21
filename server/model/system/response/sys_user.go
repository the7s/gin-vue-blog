package response

import (
	"github.com/the7s/go-vue-blog/server/global"
	"github.com/the7s/go-vue-blog/server/model/system"
)

type LoginResponse struct {
	User      system.User   `json:"user"`
	Token     string        `json:"token"`
	ExpiresAt global.MyTime `json:"expiresAt"`
}
