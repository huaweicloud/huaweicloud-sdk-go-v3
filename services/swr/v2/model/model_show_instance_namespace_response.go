package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceNamespaceResponse Response Object
type ShowInstanceNamespaceResponse struct {

	// 命名空间名称
	Name *string `json:"name,omitempty"`

	Metadata *NamespaceMetadata `json:"metadata,omitempty"`

	// 命名空间ID
	NamespaceId *int32 `json:"namespace_id,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 镜像数量
	RepoCount      *int32 `json:"repo_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowInstanceNamespaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceNamespaceResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceNamespaceResponse", string(data)}, " ")
}
