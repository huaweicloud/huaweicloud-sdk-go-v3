package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowShardingBalancerRequest struct {
	// 实例ID，可以调用“查询实例列表”接口获取。如果未申请实例，可以调用“创建实例”接口创建。

	InstanceId string `json:"instance_id"`
}

func (o ShowShardingBalancerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowShardingBalancerRequest struct{}"
	}

	return strings.Join([]string{"ShowShardingBalancerRequest", string(data)}, " ")
}
