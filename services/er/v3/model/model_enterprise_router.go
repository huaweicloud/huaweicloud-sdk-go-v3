package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnterpriseRouter 企业路由器
type EnterpriseRouter struct {

	// 企业路由器实例的ID
	Id string `json:"id"`

	// 企业路由器实例名称
	Name string `json:"name"`

	// 企业路由器实例描述信息
	Description *string `json:"description,omitempty"`

	// 运行状态
	State string `json:"state"`

	// 标签信息
	Tags *[]Tag `json:"tags,omitempty"`

	// 计费模式 按需
	ChargeMode *string `json:"charge_mode,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 项目ID
	ProjectId string `json:"project_id"`

	// 企业路由器实例的BGP AS号
	Asn int64 `json:"asn"`

	// 是否开启默认路由表传播，默认false不开启
	EnableDefaultPropagation bool `json:"enable_default_propagation"`

	// 是否开启默认路由表关联，默认false不开启
	EnableDefaultAssociation bool `json:"enable_default_association"`

	// 默认传播路由表id
	DefaultPropagationRouteTableId *string `json:"default_propagation_route_table_id,omitempty"`

	// 默认关联路由表id
	DefaultAssociationRouteTableId *string `json:"default_association_route_table_id,omitempty"`

	// 企业路由器所在可用区信息
	AvailabilityZoneIds []string `json:"availability_zone_ids"`

	// 是否自动接受共享连接创建，默认false不开启
	AutoAcceptSharedAttachments *bool `json:"auto_accept_shared_attachments,omitempty"`

	// domain id
	DomainId *string `json:"domain_id,omitempty"`
}

func (o EnterpriseRouter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnterpriseRouter struct{}"
	}

	return strings.Join([]string{"EnterpriseRouter", string(data)}, " ")
}
