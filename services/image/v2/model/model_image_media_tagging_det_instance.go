package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageMediaTaggingDetInstance struct {

	// 目标检测框位置信息，包括4个值：  width：检测框区域宽度  height：检测框区域高度  top_left_x：检测框左上角到垂直轴距离  top_left_y：检测框左上角到水平轴距离
	BoundingBox *interface{} `json:"bounding_box,omitempty"`

	// 检测标签的置信度，将Float型置信度转为String类型返回，Float取值范围（0~100）。
	Confidence *string `json:"confidence,omitempty"`
}

func (o ImageMediaTaggingDetInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageMediaTaggingDetInstance struct{}"
	}

	return strings.Join([]string{"ImageMediaTaggingDetInstance", string(data)}, " ")
}
