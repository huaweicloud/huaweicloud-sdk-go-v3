package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateSecretResponse struct {

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
	Tags           *[]Tag `json:"tags,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateSecretResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecretResponse struct{}"
	}

	return strings.Join([]string{"UpdateSecretResponse", string(data)}, " ")
}
