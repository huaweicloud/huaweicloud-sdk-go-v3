package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkspaceResponse Response Object
type CreateWorkspaceResponse struct {

	// 工作空间id
	Id *string `json:"id,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 工作空间名称
	Name *string `json:"name,omitempty"`

	// 工作空间描述
	Description *string `json:"description,omitempty"`

	// 创建人id
	CreatorId *string `json:"creator_id,omitempty"`

	// 创建人名称
	CreatorName *string `json:"creator_name,omitempty"`

	// 修改人id
	ModifierId *string `json:"modifier_id,omitempty"`

	// 修改人名称
	ModifierName *string `json:"modifier_name,omitempty"`

	// 所属项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 所属项目名称
	ProjectName *string `json:"project_name,omitempty"`

	// 所属租户id
	DomainId *string `json:"domain_id,omitempty"`

	// 所属租户名称
	DomainName *string `json:"domain_name,omitempty"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 企业项目名称
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`

	// 是否是视图
	IsView *bool `json:"is_view,omitempty"`

	// 区域id
	RegionId *string `json:"region_id,omitempty"`

	// 视图绑定的空间id
	ViewBindId *string `json:"view_bind_id,omitempty"`

	// 视图绑定的空间名称
	ViewBindName *string `json:"view_bind_name,omitempty"`

	// 仅用于视图场景，列出了该视图纳管的空间列表
	WorkspaceAgencyList *[]CreateWorkspaceResponseBodyWorkspaceAgencyList `json:"workspace_agency_list,omitempty"`
	HttpStatusCode      int                                               `json:"-"`
}

func (o CreateWorkspaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkspaceResponse struct{}"
	}

	return strings.Join([]string{"CreateWorkspaceResponse", string(data)}, " ")
}
