package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BusinessCardImageUrl 用于制作数字人名片的图片下载URL。
type BusinessCardImageUrl struct {

	// 任务照片下载URL。
	HumanImageUrl *string `json:"human_image_url,omitempty"`

	// Logo图片下载URL。
	LogoImage *string `json:"logo_image,omitempty"`
}

func (o BusinessCardImageUrl) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BusinessCardImageUrl struct{}"
	}

	return strings.Join([]string{"BusinessCardImageUrl", string(data)}, " ")
}
