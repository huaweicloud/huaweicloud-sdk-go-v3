package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 监听器对象列表，用于状态树中
type ListenersInStatusResp struct {
	// 监听器ID

	Id string `json:"id"`
	// 监听器名称

	Name string `json:"name"`
	// 监听器关联的后端云服务器组列表

	Pools []PoolsInStatusResp `json:"pools"`
	// 监听器关联的转发策略列表

	L7policies []L7policiesInStatusResp `json:"l7policies"`
	// 监听器的操作状态；该字段为预留字段，暂未启用。默认为ONLINE。

	OperatingStatus string `json:"operating_status"`
	// 监听器的配置状态；该字段为预留字段，暂未启用。默认为ACTIVE。

	ProvisioningStatus string `json:"provisioning_status"`
}

func (o ListenersInStatusResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListenersInStatusResp struct{}"
	}

	return strings.Join([]string{"ListenersInStatusResp", string(data)}, " ")
}
