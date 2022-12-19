package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 目标检测框信息。采用万分比相对坐标：即假定任何图像的宽高均为10000*10000时，其目标框的坐标值。在实际图像中恢复实际像素坐标系的方法：X*1920/10000,Y*1080/10000;
type TargetPos struct {

	// **参数说明**：目标区域框左上X坐标。
	LeftTopX *int64 `json:"left_top_x,omitempty"`

	// **参数说明**：目标区域框左上Y坐标。
	LeftTopY *int64 `json:"left_top_y,omitempty"`

	// **参数说明**：目标区域框右下X坐标。
	RightBottomX *int64 `json:"right_bottom_x,omitempty"`

	// **参数说明**：目标区域框右下Y坐标。
	RightBottomY *int64 `json:"right_bottom_y,omitempty"`
}

func (o TargetPos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TargetPos struct{}"
	}

	return strings.Join([]string{"TargetPos", string(data)}, " ")
}
