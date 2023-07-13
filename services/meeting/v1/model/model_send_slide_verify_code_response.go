package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendSlideVerifyCodeResponse Response Object
type SendSlideVerifyCodeResponse struct {

	// 抠出图形后的原背景图。通过“data:url”方式来定义图片。
	ShadowImage *string `json:"shadowImage,omitempty"`

	// 抠出的图形。
	CutImage *string `json:"cutImage,omitempty"`

	// 抠出图形的Y轴座标。
	PointY *int32 `json:"pointY,omitempty"`

	// 验证码Token字符串。
	Token *string `json:"token,omitempty"`

	// 验证码有效时间，单位：秒。
	Expire         *int32 `json:"expire,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o SendSlideVerifyCodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendSlideVerifyCodeResponse struct{}"
	}

	return strings.Join([]string{"SendSlideVerifyCodeResponse", string(data)}, " ")
}
