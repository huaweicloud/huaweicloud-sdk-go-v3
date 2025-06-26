package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Repository struct {

	// 仓库ID
	Id *int32 `json:"id,omitempty"`

	// 仓库名称
	Name *string `json:"name,omitempty"`

	// 命名空间名称
	NamespaceName *string `json:"namespace_name,omitempty"`

	// 命名空间ID
	NamespaceId *int32 `json:"namespace_id,omitempty"`

	// 仓库内的制品版本数量
	TagCount *int32 `json:"tag_count,omitempty"`

	// 被下载总次数
	PullCount *int32 `json:"pull_count,omitempty"`

	// 制品包总数
	ArtifactCount *int32 `json:"artifact_count,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o Repository) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Repository struct{}"
	}

	return strings.Join([]string{"Repository", string(data)}, " ")
}
