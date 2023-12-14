package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TopoInstanceInfo 集群实例信息
type TopoInstanceInfo struct {

	// 实例ID
	Id *string `json:"id,omitempty"`

	// 实例名称
	Name *string `json:"name,omitempty"`

	// 实例管理IP
	ManageIp *string `json:"manage_ip,omitempty"`

	// 业务IP
	TrafficIp *string `json:"traffic_ip,omitempty"`

	// 内部通信IP
	InternalIp *string `json:"internal_ip,omitempty"`

	// 内部管理IP
	InternalMgntIp *string `json:"internal_mgnt_ip,omitempty"`

	// EIP
	Eip *string `json:"eip,omitempty"`

	// elb地址
	Elb *string `json:"elb,omitempty"`

	// 实例状态
	Status *string `json:"status,omitempty"`

	// 可用区编码
	AzCode *string `json:"az_code,omitempty"`
}

func (o TopoInstanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopoInstanceInfo struct{}"
	}

	return strings.Join([]string{"TopoInstanceInfo", string(data)}, " ")
}
