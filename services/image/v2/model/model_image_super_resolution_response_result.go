package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageSuperResolutionResponseResult 图像超分结果信息
type ImageSuperResolutionResponseResult struct {

	// 超分结果图片的base64码
	ResultBase64 *string `json:"result_base64,omitempty"`
}

func (o ImageSuperResolutionResponseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageSuperResolutionResponseResult struct{}"
	}

	return strings.Join([]string{"ImageSuperResolutionResponseResult", string(data)}, " ")
}
