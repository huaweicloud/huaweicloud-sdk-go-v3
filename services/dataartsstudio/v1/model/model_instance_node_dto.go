package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceNodeDto 集群节点信息
type InstanceNodeDto struct {

	// 节点ID。
	Id *string `json:"id,omitempty"`

	// 节点名称。
	Name *string `json:"name,omitempty"`

	// 节点IP地址。
	PrivateIp *string `json:"private_ip,omitempty"`

	// 节点状态。
	Status *string `json:"status,omitempty"`

	// 节点创建人。
	CreateUser *string `json:"create_user,omitempty"`

	// 节点创建时间。
	CreateTime *string `json:"create_time,omitempty"`

	// 节点版本号。
	GatewayVersion *string `json:"gateway_version,omitempty"`
}

func (o InstanceNodeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceNodeDto struct{}"
	}

	return strings.Join([]string{"InstanceNodeDto", string(data)}, " ")
}
