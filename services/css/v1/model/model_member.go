package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Member struct {

	// 后端服务器名称。
	Name *string `json:"name,omitempty"`

	// 后端服务器对应的IP地址。
	Address *string `json:"address,omitempty"`

	// 后端服务器业务端口号。
	ProtocolPort *int32 `json:"protocol_port,omitempty"`

	// 后端云服务器的健康状态。
	OperatingStatus *string `json:"operating_status,omitempty"`

	// member关联的实例ID。
	InstanceId *string `json:"instance_id,omitempty"`
}

func (o Member) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Member struct{}"
	}

	return strings.Join([]string{"Member", string(data)}, " ")
}
