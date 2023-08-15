package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageReadConfig
type ImageReadConfig struct {

	// 播报内容，长度为1~2500
	ReadContent string `json:"read_content"`

	// 图片id
	ImageId string `json:"image_id"`

	Resolution *Resolution `json:"resolution"`
}

func (o ImageReadConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageReadConfig struct{}"
	}

	return strings.Join([]string{"ImageReadConfig", string(data)}, " ")
}
