package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 后端云服务器对象列表，用于状态树中
type PoolsInStatusResp struct {
	// 后端云服务器组ID

	Id string `json:"id"`
	// 后端云服务器组名称

	Name string `json:"name"`
	// 后端云服务器组关联的后端云服务器列表

	Members []MembersInStatusResp `json:"members"`
	// 后端云服务器组的操作状态；该字段为预留字段，暂未启用。默认为ONLINE。

	OperatingStatus string `json:"operating_status"`
	// 后端云服务器组的配置状态；该字段为预留字段，暂未启用。默认为ACTIVE。

	ProvisioningStatus string `json:"provisioning_status"`

	Healthmonitor *HealthmonitorsInStatusResp `json:"healthmonitor"`
}

func (o PoolsInStatusResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolsInStatusResp struct{}"
	}

	return strings.Join([]string{"PoolsInStatusResp", string(data)}, " ")
}
