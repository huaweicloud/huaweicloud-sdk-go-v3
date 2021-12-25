package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListResponsePropertiesRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 服务ID

	ServiceId string `json:"service_id"`
	// 命令ID

	CommandId int32 `json:"command_id"`
	// 每页显示条目数量，最大数量999，超过999后只返回999

	Limit *int32 `json:"limit,omitempty"`
	// 响应属性ID

	ResponseId *int32 `json:"response_id,omitempty"`
	// 响应属性名称

	ResponseName *string `json:"response_name,omitempty"`
	// 偏移量，表示从此偏移量开始查询， offset大于等于0

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListResponsePropertiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResponsePropertiesRequest struct{}"
	}

	return strings.Join([]string{"ListResponsePropertiesRequest", string(data)}, " ")
}
