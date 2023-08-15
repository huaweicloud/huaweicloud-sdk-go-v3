package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WorkspaceListElem struct {

	// 工作空间id
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 对应的Modelarts工作空间的id
	MaWorkspaceId *string `json:"ma_workspace_id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 工作空间名称
	Name *string `json:"name,omitempty"`

	// 工作空间描述
	Description *string `json:"description,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 企业项目名称
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 创建者
	Owner *string `json:"owner,omitempty"`

	// 创建者的USER_ID
	UserId *string `json:"user_id,omitempty"`

	// 正在被使用
	UsedFlag *string `json:"used_flag,omitempty"`
}

func (o WorkspaceListElem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkspaceListElem struct{}"
	}

	return strings.Join([]string{"WorkspaceListElem", string(data)}, " ")
}
