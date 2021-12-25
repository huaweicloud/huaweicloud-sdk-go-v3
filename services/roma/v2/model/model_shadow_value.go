package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShadowValue struct {
	// 属性名称

	PropertyName *string `json:"property_name,omitempty"`
	// 属性最后一次上报值

	PropertyValue *string `json:"property_value,omitempty"`
	// 属性最后一次上报时间，格式timestamp(ms)，使用UTC时区

	PropertyUpdatedDate *int64 `json:"property_updated_date,omitempty"`
}

func (o ShadowValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShadowValue struct{}"
	}

	return strings.Join([]string{"ShadowValue", string(data)}, " ")
}
