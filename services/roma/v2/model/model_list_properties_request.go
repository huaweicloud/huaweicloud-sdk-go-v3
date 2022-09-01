package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPropertiesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 服务ID
	ServiceId string `json:"service_id" xml:"service_id"`

	// 每页显示条目数量，最大数量999，超过999后只返回999
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 属性ID
	PropertyId *int32 `json:"property_id,omitempty" xml:"property_id"`

	// 属性名称
	PropertyName *string `json:"property_name,omitempty" xml:"property_name"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0
	Offset *int32 `json:"offset,omitempty" xml:"offset"`
}

func (o ListPropertiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPropertiesRequest struct{}"
	}

	return strings.Join([]string{"ListPropertiesRequest", string(data)}, " ")
}
