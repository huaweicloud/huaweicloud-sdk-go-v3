package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageMainObjectDetectionInstance struct {

	// 主体的类别，现阶段分为：bounding_box和main_object_box
	Label *string `json:"label,omitempty"`

	//  目标检测框位置信息，包括4个值：  width：检测框区域宽度  height：检测框区域高度  top_left_x：检测框左上角到垂直轴距离  top_left_y：检测框左上角到水平轴距离 properties: width: type: string description: 检测框区域高度 example: 139.58 height: type: string description: 检测框区域高度 example: 261.32 top_left_x: type: string description: 检测框左上角到垂直轴距离 example: 256.13 top_left_y: type: string description: 检测框左上角到水平轴距离 example: 85.2
	Location *interface{} `json:"location,omitempty"`

	// 主体框的置信度,将Float型置信度转为String类型返回,Float取值范围（0~100）。
	Confidence *string `json:"confidence,omitempty"`
}

func (o ImageMainObjectDetectionInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageMainObjectDetectionInstance struct{}"
	}

	return strings.Join([]string{"ImageMainObjectDetectionInstance", string(data)}, " ")
}
