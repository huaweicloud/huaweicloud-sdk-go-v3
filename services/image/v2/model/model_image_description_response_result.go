package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 图像描述结果信息
type ImageDescriptionResponseResult struct {

	// 图像描述。
	Description *string `json:"description,omitempty"`
}

func (o ImageDescriptionResponseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageDescriptionResponseResult struct{}"
	}

	return strings.Join([]string{"ImageDescriptionResponseResult", string(data)}, " ")
}
