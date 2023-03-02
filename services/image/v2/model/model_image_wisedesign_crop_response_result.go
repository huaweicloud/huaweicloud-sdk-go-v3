package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 智能设计商品图像裁剪结果信息
type ImageWisedesignCropResponseResult struct {

	// 裁剪结果框，有四个整型数值，分别代表左上角横坐标、左上角纵坐标、右下角横坐标、右下角纵坐标，示例：[x_min, y_min, x_max, y_max]
	Box *[]int32 `json:"box,omitempty"`

	// 裁剪后图像的64位编码
	ImageBase64 *string `json:"image_base64,omitempty"`

	// 请求的任务id
	TaskId *string `json:"task_id,omitempty"`
}

func (o ImageWisedesignCropResponseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageWisedesignCropResponseResult struct{}"
	}

	return strings.Join([]string{"ImageWisedesignCropResponseResult", string(data)}, " ")
}
