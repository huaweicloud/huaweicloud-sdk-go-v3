package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageMediaTaggingDetInstance struct {
	BoundingBox *BoundingBox `json:"bounding_box,omitempty"`

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
