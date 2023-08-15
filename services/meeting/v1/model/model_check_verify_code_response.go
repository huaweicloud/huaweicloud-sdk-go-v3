package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckVerifyCodeResponse Response Object
type CheckVerifyCodeResponse struct {

	// 访问token字符串。
	Token *string `json:"token,omitempty"`

	// 过期时间，单位：秒。
	Expire         *int32 `json:"expire,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CheckVerifyCodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckVerifyCodeResponse struct{}"
	}

	return strings.Join([]string{"CheckVerifyCodeResponse", string(data)}, " ")
}
