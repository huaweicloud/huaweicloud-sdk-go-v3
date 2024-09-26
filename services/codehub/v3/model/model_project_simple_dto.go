package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProjectSimpleDto struct {

	// 项目id
	Id *int32 `json:"id,omitempty"`

	// 项目描述
	Description *string `json:"description,omitempty"`

	// 项目名称
	Name *string `json:"name,omitempty"`

	// 项目名称
	NameWithNamespace *string `json:"name_with_namespace,omitempty"`

	// 项目路径
	Path *string `json:"path,omitempty"`

	// 项目路径
	PathWithNamespace *string `json:"path_with_namespace,omitempty"`

	// 开发模式
	DevelopMode *string `json:"develop_mode,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 是否归档
	Archived *bool `json:"archived,omitempty"`

	// 是否为kia仓
	IsKia *bool `json:"is_kia,omitempty"`

	// 项目url
	SshUrlToRepo *string `json:"ssh_url_to_repo,omitempty"`

	// 项目url
	HttpUrlToRepo *string `json:"http_url_to_repo,omitempty"`

	// 项目url
	WebUrl *string `json:"web_url,omitempty"`

	// 项目readme url
	ReadmeUrl *string `json:"readme_url,omitempty"`

	// product id
	ProductId *string `json:"product_id,omitempty"`

	// product name
	ProductName *string `json:"product_name,omitempty"`

	// member mgnt mode
	MemberMgntMode *int32 `json:"member_mgnt_mode,omitempty"`

	// visibility
	Visibility *string `json:"visibility,omitempty"`

	Namespace *NamespaceBasicDto `json:"namespace,omitempty"`

	// 项目类型
	RepoType *string `json:"repo_type,omitempty"`
}

func (o ProjectSimpleDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectSimpleDto struct{}"
	}

	return strings.Join([]string{"ProjectSimpleDto", string(data)}, " ")
}
