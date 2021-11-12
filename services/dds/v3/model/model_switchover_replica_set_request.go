package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SwitchoverReplicaSetRequest struct {
	// 实例ID，可以调用“查询实例列表”接口获取。如果未申请实例，可以调用“创建实例”接口创建。

	InstanceId string `json:"instance_id"`
}

func (o SwitchoverReplicaSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchoverReplicaSetRequest struct{}"
	}

	return strings.Join([]string{"SwitchoverReplicaSetRequest", string(data)}, " ")
}
