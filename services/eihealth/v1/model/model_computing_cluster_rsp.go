package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ComputingClusterRsp 计算集群详情。
type ComputingClusterRsp struct {

	// 已绑定的集群ID。
	Id *string `json:"id,omitempty"`

	// 计算集群名称。
	Name *string `json:"name,omitempty"`

	// 计算集群版本。
	Version *string `json:"version,omitempty"`

	// 计算集群规格。
	Flavor *string `json:"flavor,omitempty"`

	// 计算集群类别。CCE代表CCE集群，Turbo代表CCE Turbo集群。
	Category *string `json:"category,omitempty"`

	// 计算集群绑定状态，支持RUNNING|INSTALL_FAILED|UNINSTALL_FAILED|INSTALLING|UNINSTALLING。
	Status *string `json:"status,omitempty"`

	// 计算集群可用节点数。
	ActiveNodesNumber *int32 `json:"active_nodes_number,omitempty"`

	// 计算集群总节点数。
	TotalNodesNumber *int32 `json:"total_nodes_number,omitempty"`

	// 计算集群绑定时间。
	CreateTime *string `json:"create_time,omitempty"`

	// cce集群创建时间。
	CceCreateTime *string `json:"cce_create_time,omitempty"`

	// cce集群绑态。
	CceStatus *string `json:"cce_status,omitempty"`

	// cce集群ID。
	CceId *string `json:"cce_id,omitempty"`

	// 计费模式，prepaid表示包周期，postpaid表示按需。
	BillingMode *string `json:"billing_mode,omitempty"`
}

func (o ComputingClusterRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComputingClusterRsp struct{}"
	}

	return strings.Join([]string{"ComputingClusterRsp", string(data)}, " ")
}
