package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GroupDto A group object that contains a specified group’s metadata and attributes.
type GroupDto struct {

	// 包含用户组描述的字符串
	Description *string `json:"description,omitempty"`

	// 包含用户组显示名称的字符串
	DisplayName *string `json:"display_name,omitempty"`

	// 外部身份源分配给此资源的标识符
	ExternalId *string `json:"external_id,omitempty"`

	// 包含外部身份提供商颁发给此资源的标识符的对象列表
	ExternalIds *[]ExternalIdDto `json:"external_ids,omitempty"`

	// 身份源中IdentityCenter用户组的全局唯一标识符（ID）
	GroupId string `json:"group_id"`

	// 身份源的全局唯一标识符（ID）
	IdentityStoreId string `json:"identity_store_id"`

	// 创建时的时间戳
	CreatedAt *int64 `json:"created_at,omitempty"`

	// 创建者
	CreatedBy *string `json:"created_by,omitempty"`

	// 更新时的时间戳
	UpdatedAt *int64 `json:"updated_at,omitempty"`

	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty"`
}

func (o GroupDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupDto struct{}"
	}

	return strings.Join([]string{"GroupDto", string(data)}, " ")
}
