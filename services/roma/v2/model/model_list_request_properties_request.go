package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListRequestPropertiesRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 服务ID

	ServiceId string `json:"service_id"`
	// 命令ID

	CommandId int32 `json:"command_id"`
	// 每页显示条目数量，最大数量999，超过999后只返回999

	Limit *int32 `json:"limit,omitempty"`
	// 请求属性ID

	RequestId *int32 `json:"request_id,omitempty"`
	// 请求属性名称

	RequestName *string `json:"request_name,omitempty"`
	// 偏移量，表示从此偏移量开始查询， offset大于等于0

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListRequestPropertiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRequestPropertiesRequest struct{}"
	}

	return strings.Join([]string{"ListRequestPropertiesRequest", string(data)}, " ")
}
