package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssetPropertyLastValue struct {

	// 资产属性名称
	PropertyName *string `json:"property_name,omitempty" xml:"property_name"`

	// 资产属性值
	Value *interface{} `json:"value,omitempty" xml:"value"`

	// 资产属性值最后更新时间
	Timestamp *string `json:"timestamp,omitempty" xml:"timestamp"`
}

func (o AssetPropertyLastValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetPropertyLastValue struct{}"
	}

	return strings.Join([]string{"AssetPropertyLastValue", string(data)}, " ")
}
