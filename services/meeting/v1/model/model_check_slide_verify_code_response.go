package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CheckSlideVerifyCodeResponse struct {

	// 访问token字符串
	Token *string `json:"token,omitempty" xml:"token"`

	// 过期时间，单位：秒
	Expire         *int32 `json:"expire,omitempty" xml:"expire"`
	HttpStatusCode int    `json:"-"`
}

func (o CheckSlideVerifyCodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckSlideVerifyCodeResponse struct{}"
	}

	return strings.Join([]string{"CheckSlideVerifyCodeResponse", string(data)}, " ")
}
