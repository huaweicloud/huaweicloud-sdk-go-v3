package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateProjectRepoRequest struct {

	// 是否导入项目成员，取值范围：0->不导入项目成员，1->导入项目成员
	ImportMembers *int32 `json:"import_members,omitempty"`

	// 项目名称，取值范围：可以输入英文大小写字母、数字、连字符、下划线，且必须以字母开头
	ProjectName string `json:"project_name"`

	// 仓库名称，取值范围：可以输入英文大小写字母、数字、连字符、下划线，且必须以字母开头
	RepoName string `json:"repo_name"`

	// 项目类型，normal|scrum
	Type *string `json:"type,omitempty"`

	// 是否是公仓
	VisibilityLevel *int32 `json:"visibility_level,omitempty"`

	ExternalProjectInfo *ExternalKeyMessage `json:"external_project_info,omitempty"`
}

func (o CreateProjectRepoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProjectRepoRequest struct{}"
	}

	return strings.Join([]string{"CreateProjectRepoRequest", string(data)}, " ")
}
