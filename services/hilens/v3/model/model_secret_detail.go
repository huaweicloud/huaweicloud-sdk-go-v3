package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecretDetail 查询密钥详情响应
type SecretDetail struct {

	// 密钥ID
	Id *string `json:"id,omitempty"`

	// 密钥名称
	Name *string `json:"name,omitempty"`

	// 密钥描述
	Description *string `json:"description,omitempty"`

	// 工作空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 密钥创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 密钥更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 密钥列表
	Secrets *[]Secret `json:"secrets,omitempty"`

	// 标签列表
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o SecretDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecretDetail struct{}"
	}

	return strings.Join([]string{"SecretDetail", string(data)}, " ")
}
