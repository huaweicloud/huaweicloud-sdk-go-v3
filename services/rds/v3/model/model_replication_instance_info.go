package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReplicationInstanceInfo 实例详情。
type ReplicationInstanceInfo struct {

	// 实例id
	InstanceId string `json:"instance_id"`

	// 实例名称。
	InstanceName string `json:"instance_name"`

	// 内网地址。
	DataVip string `json:"data_vip"`
}

func (o ReplicationInstanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReplicationInstanceInfo struct{}"
	}

	return strings.Join([]string{"ReplicationInstanceInfo", string(data)}, " ")
}
