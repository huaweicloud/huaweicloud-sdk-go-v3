package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MysqlProxyNodeV3 struct {

	// 节点ID。
	Id *string `json:"id,omitempty"`

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 节点状态。
	Status *string `json:"status,omitempty"`

	// 节点名称。
	Name *string `json:"name,omitempty"`

	// 节点读写分离读权重。
	Weight *int32 `json:"weight,omitempty"`

	// 主实例所有节点的可用区信息。
	AvailabilityZone *[]MysqlProxyAvailable `json:"availability_zone,omitempty"`
}

func (o MysqlProxyNodeV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlProxyNodeV3 struct{}"
	}

	return strings.Join([]string{"MysqlProxyNodeV3", string(data)}, " ")
}
