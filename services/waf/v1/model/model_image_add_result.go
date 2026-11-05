package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageAddResult struct {

	// 图片ID
	ImageId *string `json:"image_id,omitempty"`

	// 是否加入成功
	Success *bool `json:"success,omitempty"`

	// 失败原因（如图片id不存在/大小超限/格式不支持）
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o ImageAddResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageAddResult struct{}"
	}

	return strings.Join([]string{"ImageAddResult", string(data)}, " ")
}
