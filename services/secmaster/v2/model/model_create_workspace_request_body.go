package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkspaceRequestBody 创建工作空间时参数对象
type CreateWorkspaceRequestBody struct {

	// 区域id
	RegionId string `json:"region_id"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 企业项目名称
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`

	// 视图绑定的空间id
	ViewBindId *string `json:"view_bind_id,omitempty"`

	// 是否是视图
	IsView *bool `json:"is_view,omitempty"`

	// 工作空间名称
	Name string `json:"name"`

	// 工作空间描述
	Description *string `json:"description,omitempty"`

	// 项目名称
	ProjectName string `json:"project_name"`

	// 标签数组
	Tags *[]TagsPojo `json:"tags,omitempty"`
}

func (o CreateWorkspaceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkspaceRequestBody struct{}"
	}

	return strings.Join([]string{"CreateWorkspaceRequestBody", string(data)}, " ")
}
