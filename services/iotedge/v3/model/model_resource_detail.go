package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceDetail 资源信息。
type ResourceDetail struct {

	// 资源id，添加资源时由边缘侧生成
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源状态，冻结:freeze、解冻:unfreeze、退订:delete。
	Status *string `json:"status,omitempty"`

	// 计费规则
	ChargingRule *string `json:"charging_rule,omitempty"`

	// 内部类型
	Type *string `json:"type,omitempty"`

	// 资源名称，由边缘侧生成。
	ResourceName *string `json:"resource_name,omitempty"`

	// 公有云CBC上注册的服务类型英文名
	CloudServiceType *string `json:"cloud_service_type,omitempty"`

	// CBC上注册的资源类型编码。
	ResourceType *string `json:"resource_type,omitempty"`

	// 资源规格编码
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// 关联的边缘集群ID
	AssociatedEdgeClusterId *string `json:"associated_edge_cluster_id,omitempty"`
}

func (o ResourceDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceDetail struct{}"
	}

	return strings.Join([]string{"ResourceDetail", string(data)}, " ")
}
