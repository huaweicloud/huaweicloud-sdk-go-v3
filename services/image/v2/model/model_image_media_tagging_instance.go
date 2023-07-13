package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageMediaTaggingInstance struct {
	BoundingBox *BoundingBox `json:"bounding_box,omitempty"`

	// 检测标签置信度,将Float型置信度转为String类型返回,Float取值范围（0~100）。
	Confidence *string `json:"confidence,omitempty"`
}

func (o ImageMediaTaggingInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageMediaTaggingInstance struct{}"
	}

	return strings.Join([]string{"ImageMediaTaggingInstance", string(data)}, " ")
}
