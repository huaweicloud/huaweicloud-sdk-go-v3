package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PublisherVo struct {

	// **参数解释**： 发布商ID。可通过[查询发布商详情](ShowPublisher.xml)查询。 **取值范围**： 不涉及。
	PublisherUniqueId *string `json:"publisher_unique_id,omitempty"`

	// **参数解释**： 用户ID。 **取值范围**： 不涉及。
	UserId *string `json:"user_id,omitempty"`

	// **参数解释**： 租户ID。 **取值范围**： 不涉及。
	TenantId *string `json:"tenant_id,omitempty"`

	// **参数解释**： 发布商描述。 **取值范围**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 发布商图标URL。 **取值范围**： 不涉及。
	LogoUrl *string `json:"logo_url,omitempty"`

	// **参数解释**： 创建发布商时填入的网页地址。 **取值范围**： 不涉及。
	Website *string `json:"website,omitempty"`

	// **参数解释**： 创建发布商时填入的帮助地址。 **取值范围**： 不涉及。
	SupportUrl *string `json:"support_url,omitempty"`

	// **参数解释**： 创建发布商时填入的源地址。 **取值范围**： 不涉及。
	SourceUrl *string `json:"source_url,omitempty"`

	// **参数解释**： 发布商英文名。 **取值范围**： 不涉及。
	EnName *string `json:"en_name,omitempty"`

	// **参数解释**： 发布商中文名。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 授权状态。 **取值范围**： 不涉及。
	AuthStatus *string `json:"auth_status,omitempty"`

	// **参数解释**： 是否删除。 **取值范围**： - 0：未删除。 - 1：已删除。
	IsDelete *int32 `json:"is_delete,omitempty"`

	// **参数解释**： 最后更新人名称。 **取值范围**： 不涉及。
	LastUpdateUserName *string `json:"last_update_user_name,omitempty"`

	// **参数解释**： 最后更新人用户ID。 **取值范围**： 不涉及。
	LastUpdateUserId *string `json:"last_update_user_id,omitempty"`

	// **参数解释**： 最后更新时间。 **取值范围**： 不涉及。
	LastUpdateTime *string `json:"last_update_time,omitempty"`
}

func (o PublisherVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublisherVo struct{}"
	}

	return strings.Join([]string{"PublisherVo", string(data)}, " ")
}
