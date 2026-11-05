package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageRemoveResult struct {

	// 图片ID
	ImageId *string `json:"image_id,omitempty"`

	// 是否移除成功
	Success *bool `json:"success,omitempty"`

	// 失败原因
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o ImageRemoveResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageRemoveResult struct{}"
	}

	return strings.Join([]string{"ImageRemoveResult", string(data)}, " ")
}
