package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 智能设计图像修复结果信息
type ImageWisedesignInpaintingResponseResult struct {

	// 修复结果图片base64
	ImageBase64 *string `json:"image_base64,omitempty"`

	// 请求类型
	Action *string `json:"action,omitempty"`

	// 请求的任务id
	TaskId *string `json:"task_id,omitempty"`
}

func (o ImageWisedesignInpaintingResponseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageWisedesignInpaintingResponseResult struct{}"
	}

	return strings.Join([]string{"ImageWisedesignInpaintingResponseResult", string(data)}, " ")
}
