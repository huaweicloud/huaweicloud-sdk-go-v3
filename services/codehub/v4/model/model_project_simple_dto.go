package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProjectSimpleDto Get the simplest project info, used to the list of projects
type ProjectSimpleDto struct {

	// 仓库id
	Id *int32 `json:"id,omitempty"`

	// 仓库描述信息
	Description *string `json:"description,omitempty"`

	// 仓库名称
	Name *string `json:"name,omitempty"`

	// 组织名/组织名.../仓库名
	NameWithNamespace *string `json:"name_with_namespace,omitempty"`

	// 仓库路径
	Path *string `json:"path,omitempty"`

	// 仓库完整路径
	PathWithNamespace *string `json:"path_with_namespace,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 是否归档
	Archived *bool `json:"archived,omitempty"`

	// 是否机密
	IsKia *bool `json:"is_kia,omitempty"`

	// 仓库ssh地址
	SshUrlToRepo *string `json:"ssh_url_to_repo,omitempty"`

	// 仓库http地址
	HttpUrlToRepo *string `json:"http_url_to_repo,omitempty"`

	// 仓库页面链接
	WebUrl *string `json:"web_url,omitempty"`

	// 仓库readme文件链接
	ReadmeUrl *string `json:"readme_url,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 项目名称
	ProjectName *string `json:"project_name,omitempty"`

	// 仓库开发模式：normal、CR
	DevelopMode *string `json:"develop_mode,omitempty"`

	// 审核状态
	ModerationResult *bool `json:"moderation_result,omitempty"`
}

func (o ProjectSimpleDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectSimpleDto struct{}"
	}

	return strings.Join([]string{"ProjectSimpleDto", string(data)}, " ")
}
