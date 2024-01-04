package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ForkProjectRepoRequest struct {

	// 是否导入项目成员，取值范围：0->不导入项目成员，1->导入项目成员
	ImportMembers *int32 `json:"import_members,omitempty"`

	// 项目名称，取值范围：可以输入英文大小写字母、数字、连字符、下划线，且必须以字母开头
	ProjectName string `json:"project_name"`

	// 仓库名称，取值范围：可以输入英文大小写字母、数字、连字符、下划线，且必须以字母开头
	RepoName string `json:"repo_name"`

	// 复制模板的ID
	TemplateId string `json:"template_id"`

	// 项目类型，scrum
	Type *string `json:"type,omitempty"`

	// 仓库可见性：  *私有仓库：仓库仅对仓库成员可见，仓库成员可读写和访问仓库，取值范围为0  *公开仓库：   1.项目内成员只读仓库：仓库对项目内成员公开只读，并项目内成员可在项目下和代码组下的仓库列表中查看和搜索，取值范围为10   2.租户内成员只读仓库：仓库对租户内成员公开只读，并租户内成员可在项目下和代码组下的仓库列表中查看和搜索，取值范围为10   3.所有访客只读仓库：仓库对所有访客公开只读，并所有访客可在项目下和代码组下的仓库列表中查看和搜索，取值范围为20
	VisibilityLevel *int32 `json:"visibility_level,omitempty"`

	ExternalProjectInfo *ExternalKeyMessage `json:"external_project_info,omitempty"`
}

func (o ForkProjectRepoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ForkProjectRepoRequest struct{}"
	}

	return strings.Join([]string{"ForkProjectRepoRequest", string(data)}, " ")
}
