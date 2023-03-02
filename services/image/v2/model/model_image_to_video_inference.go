package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 图像合成视频服务推理参数
type ImageToVideoInference struct {
	ImageConfig *ImageToVideoInfo `json:"image_config"`
}

func (o ImageToVideoInference) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageToVideoInference struct{}"
	}

	return strings.Join([]string{"ImageToVideoInference", string(data)}, " ")
}
