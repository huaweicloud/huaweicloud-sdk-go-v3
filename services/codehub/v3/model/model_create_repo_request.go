package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRepoRequest struct {

	// 是否导入项目成员，取值范围：0->不导入项目成员，1->导入项目成员
	ImportMembers *int32 `json:"import_members,omitempty"`

	// 仓库名称，取值范围：可以输入英文大小写字母、数字、连字符、下划线，且必须以字母开头
	Name string `json:"name"`

	// 指定项目的UUID
	ProjectUuid string `json:"project_uuid"`

	// 复制模板的ID
	TemplateId *string `json:"template_id,omitempty"`

	// 仓库状态，取值范围：0->私有，20->公开只读
	VisibilityLevel *int32 `json:"visibility_level,omitempty"`

	// 模板仓库的https地址的base64加密
	ImportUrl *string `json:"import_url,omitempty"`

	// 仓库描述信息
	Description *string `json:"description,omitempty"`

	// 根据编程语言生成.gitignore文件
	GitignoreId *string `json:"gitignore_id,omitempty"`

	// 许可证id
	LicenseId *int32 `json:"license_id,omitempty"`

	// 是否允许生成README文件
	EnableReadme *int32 `json:"enable_readme,omitempty"`

	// 调用者
	Caller *string `json:"caller,omitempty"`
}

func (o CreateRepoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRepoRequest struct{}"
	}

	return strings.Join([]string{"CreateRepoRequest", string(data)}, " ")
}
