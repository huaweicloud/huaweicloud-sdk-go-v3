package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PublisherVo struct {

	// 发布商ID
	PublisherUniqueId *string `json:"publisher_unique_id,omitempty"`

	// 用户ID
	UserId *string `json:"user_id,omitempty"`

	// 租户ID
	TenantId *string `json:"tenant_id,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 图标URL
	LogoUrl *string `json:"logo_url,omitempty"`

	// 网页地址
	Website *string `json:"website,omitempty"`

	// 地址
	SupportUrl *string `json:"support_url,omitempty"`

	// 地址
	SourceUrl *string `json:"source_url,omitempty"`

	// 英文名
	EnName *string `json:"en_name,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 授权状态
	AuthStatus *string `json:"auth_status,omitempty"`

	// 是否删除
	IsDelete *int32 `json:"is_delete,omitempty"`

	// 最后更新人
	LastUpdateUserName *string `json:"last_update_user_name,omitempty"`

	// 最后更新人ID
	LastUpdateUserId *string `json:"last_update_user_id,omitempty"`

	// 最后更新时间
	LastUpdateTime *string `json:"last_update_time,omitempty"`
}

func (o PublisherVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublisherVo struct{}"
	}

	return strings.Join([]string{"PublisherVo", string(data)}, " ")
}
