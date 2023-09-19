package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BusinessCardImageConfig 用户上传的用于制作数字人名片的图片。
type BusinessCardImageConfig struct {

	// 人物照片，需要Base64编码。
	HumanImage string `json:"human_image"`

	// Logo图片，需要Base64编码。
	LogoImage *string `json:"logo_image,omitempty"`
}

func (o BusinessCardImageConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BusinessCardImageConfig struct{}"
	}

	return strings.Join([]string{"BusinessCardImageConfig", string(data)}, " ")
}
