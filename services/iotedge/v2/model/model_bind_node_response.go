package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindNodeResponse Response Object
type BindNodeResponse struct {

	// 资源ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源类型：industry|campus
	Type *string `json:"type,omitempty"`

	// 对接的子系统数量
	SubsystemCount *int32 `json:"subsystem_count,omitempty"`

	// CBC上注册的资源类型编码。
	ResourceType *string `json:"resource_type,omitempty"`

	// CBC上注册的资源类型编码。
	ResourceSpecType *string `json:"resource_spec_type,omitempty"`

	// 关联的边缘节点ID
	AssociatedEdgeNodeId *string `json:"associated_edge_node_id,omitempty"`

	// 关联的边缘节点名称
	AssociatedEdgeNodeName *string `json:"associated_edge_node_name,omitempty"`

	// 扩展开通参数。
	ExtendParams *string `json:"extend_params,omitempty"`

	// 资源容量大小，线性产品使用
	ResourceSize   *int32 `json:"resource_size,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BindNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindNodeResponse struct{}"
	}

	return strings.Join([]string{"BindNodeResponse", string(data)}, " ")
}
