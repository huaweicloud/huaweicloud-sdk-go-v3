package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEnterpriseRouter 更新企业路由器请求体
type UpdateEnterpriseRouter struct {

	// 企业路由器实例名称
	Name *string `json:"name,omitempty"`

	// 企业路由器实例描述信息
	Description *string `json:"description,omitempty"`

	// 是否开启默认传播
	EnableDefaultPropagation *bool `json:"enable_default_propagation,omitempty"`

	// 是否开启默认关联
	EnableDefaultAssociation *bool `json:"enable_default_association,omitempty"`

	// 默认传播路由表ID
	DefaultPropagationRouteTableId *string `json:"default_propagation_route_table_id,omitempty"`

	// 默认关联路由表ID
	DefaultAssociationRouteTableId *string `json:"default_association_route_table_id,omitempty"`

	// 是否自动接受共享连接创建，默认false不开启
	AutoAcceptSharedAttachments *bool `json:"auto_accept_shared_attachments,omitempty"`
}

func (o UpdateEnterpriseRouter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEnterpriseRouter struct{}"
	}

	return strings.Join([]string{"UpdateEnterpriseRouter", string(data)}, " ")
}
