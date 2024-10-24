package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UniqueAttributeDto An entity attribute that's unique to a specific entity.
type UniqueAttributeDto struct {

	// 属性路径
	AttributePath string `json:"attribute_path"`

	// 属性的值
	AttributeValue string `json:"attribute_value"`
}

func (o UniqueAttributeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UniqueAttributeDto struct{}"
	}

	return strings.Join([]string{"UniqueAttributeDto", string(data)}, " ")
}
