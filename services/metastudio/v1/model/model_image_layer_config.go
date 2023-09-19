package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageLayerConfig 素材图片图层配置。
type ImageLayerConfig struct {

	// 图片文件的URL。
	ImageUrl *string `json:"image_url,omitempty"`
}

func (o ImageLayerConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageLayerConfig struct{}"
	}

	return strings.Join([]string{"ImageLayerConfig", string(data)}, " ")
}
