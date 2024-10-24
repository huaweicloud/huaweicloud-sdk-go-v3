package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DescribeGroupResponse Response Object
type DescribeGroupResponse struct {

	// 包含组描述的字符串
	Description *string `json:"description,omitempty"`

	// 包含组显示名称的字符串
	DisplayName *string `json:"display_name,omitempty"`

	// 外部身份源分配给此资源的标识符
	ExternalId *string `json:"external_id,omitempty"`

	// 包含外部身份提供商颁发给此资源的标识符的对象列表
	ExternalIds *[]ExternalIdDto `json:"external_ids,omitempty"`

	// 身份源中IAM身份中心用户组的全局唯一标识符（ID）
	GroupId *string `json:"group_id,omitempty"`

	// 身份源的全局唯一标识符（ID）
	IdentityStoreId *string `json:"identity_store_id,omitempty"`

	// 创建时的时间戳
	CreatedAt *int64 `json:"created_at,omitempty"`

	// 创建者
	CreatedBy *string `json:"created_by,omitempty"`

	// 更新时的时间戳
	UpdatedAt *int64 `json:"updated_at,omitempty"`

	// 更新者
	UpdatedBy      *string `json:"updated_by,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DescribeGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DescribeGroupResponse struct{}"
	}

	return strings.Join([]string{"DescribeGroupResponse", string(data)}, " ")
}
