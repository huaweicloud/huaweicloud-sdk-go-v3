package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 智能设计图像滤镜结果信息
type ImageWisedesignColorfilterResponseResult struct {

	// 图片加滤镜后图像的64位编码
	ResultBase64 *string `json:"result_base64,omitempty"`
}

func (o ImageWisedesignColorfilterResponseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageWisedesignColorfilterResponseResult struct{}"
	}

	return strings.Join([]string{"ImageWisedesignColorfilterResponseResult", string(data)}, " ")
}
