package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckSlideVerifyCodeResponse Response Object
type CheckSlideVerifyCodeResponse struct {

	// 访问Token字符串。
	Token *string `json:"token,omitempty"`

	// 过期时间，单位：秒。
	Expire         *int32 `json:"expire,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CheckSlideVerifyCodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckSlideVerifyCodeResponse struct{}"
	}

	return strings.Join([]string{"CheckSlideVerifyCodeResponse", string(data)}, " ")
}
