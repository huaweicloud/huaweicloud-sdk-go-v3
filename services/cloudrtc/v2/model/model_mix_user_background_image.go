package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 指定用户的背景图，优先级大于default_user_background_image
type MixUserBackgroundImage struct {

	// 需要设置背景图的用户ID
	UserId string `json:"user_id" xml:"user_id"`

	// 需要设置背景图的地址，图片先上传obs，格式s3://bucket/object
	ImageUri string `json:"image_uri" xml:"image_uri"`
}

func (o MixUserBackgroundImage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MixUserBackgroundImage struct{}"
	}

	return strings.Join([]string{"MixUserBackgroundImage", string(data)}, " ")
}
