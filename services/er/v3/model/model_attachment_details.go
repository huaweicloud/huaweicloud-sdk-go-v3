package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AttachmentDetails struct {

	// 连接ID
	Id string `json:"id"`

	// 连接名字
	Name string `json:"name"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 连接状态:pending|available|modifying|deleting|deleted|failed|pending_acceptance|rejected|initiating_request
	State *string `json:"state,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// 企业路由器关联tag
	Tags *[]Tag `json:"tags,omitempty"`

	// 项目ID
	ProjectId string `json:"project_id"`

	// er id
	ErId *string `json:"er_id,omitempty"`

	// 内部连接关联的资源ID
	ResourceId string `json:"resource_id"`

	// - vgw：云专线的虚拟网关 - vpn：vpn网关 - gdgw：下一代专线网关 - peering：对等连接，通过云连接CC加载不同区域的企业路由器来创建“对等连接（Peering）”连接 - can：智能云接入网关
	ResourceType string `json:"resource_type"`

	// 资源所属项目ID
	ResourceProjectId *string `json:"resource_project_id,omitempty"`

	// 表示此连接是否被关联
	Associated *bool `json:"associated,omitempty"`

	// 关联路由表id
	RouteTableId *string `json:"route_table_id,omitempty"`
}

func (o AttachmentDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachmentDetails struct{}"
	}

	return strings.Join([]string{"AttachmentDetails", string(data)}, " ")
}
