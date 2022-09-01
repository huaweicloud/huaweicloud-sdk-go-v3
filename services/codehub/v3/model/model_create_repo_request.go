package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRepoRequest struct {

	// 是否导入项目成员，取值范围：0->不导入项目成员，1->导入项目成员
	ImportMembers *int32 `json:"import_members,omitempty" xml:"import_members"`

	// 仓库名称，取值范围：可以输入英文大小写字母、数字、连字符、下划线，且必须以字母开头
	Name string `json:"name" xml:"name"`

	// 指定项目的UUID
	ProjectUuid string `json:"project_uuid" xml:"project_uuid"`

	// 复制模板的ID
	TemplateId *string `json:"template_id,omitempty" xml:"template_id"`

	// 仓库状态，取值范围：0->私有，20->公开只读
	VisibilityLevel *int32 `json:"visibility_level,omitempty" xml:"visibility_level"`

	// 模板仓库的https地址的base64加密
	ImportUrl *string `json:"import_url,omitempty" xml:"import_url"`

	// 仓库描述信息
	Description *string `json:"description,omitempty" xml:"description"`

	// 根据编程语言生成.gitignore文件
	GitignoreId *string `json:"gitignore_id,omitempty" xml:"gitignore_id"`

	// 许可证id
	LicenseId *int32 `json:"license_id,omitempty" xml:"license_id"`

	// 是否允许生成README文件
	EnableReadme *int32 `json:"enable_readme,omitempty" xml:"enable_readme"`

	// 调用者
	Caller *string `json:"caller,omitempty" xml:"caller"`
}

func (o CreateRepoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRepoRequest struct{}"
	}

	return strings.Join([]string{"CreateRepoRequest", string(data)}, " ")
}
