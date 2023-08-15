package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigMap struct {

	// 配置列表
	Configs []Config `json:"configs"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 配置项ID
	Id *string `json:"id,omitempty"`

	// 名称
	Name string `json:"name"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 标签列表
	Tags []Tag `json:"tags"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 工作空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`
}

func (o ConfigMap) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigMap struct{}"
	}

	return strings.Join([]string{"ConfigMap", string(data)}, " ")
}
