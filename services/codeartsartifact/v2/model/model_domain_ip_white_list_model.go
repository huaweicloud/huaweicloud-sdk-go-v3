package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DomainIpWhiteListModel struct {

	// **参数解释**: id。 **取值范围**: 不涉及。
	Id *int64 `json:"id,omitempty"`

	// **参数解释**: 租户id。 **取值范围**: 不涉及。
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释**: 区域。 **取值范围**: 不涉及。
	Region *string `json:"region,omitempty"`

	// **参数解释**: 类型。 **取值范围**: 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**: ip地址。 **取值范围**: 不涉及。
	Value *string `json:"value,omitempty"`

	// **参数解释**: 创建时间。 **取值范围**: 不涉及。
	Created *int64 `json:"created,omitempty"`

	// **参数解释**: 更新时间。 **取值范围**: 不涉及。
	Updated *int64 `json:"updated,omitempty"`
}

func (o DomainIpWhiteListModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainIpWhiteListModel struct{}"
	}

	return strings.Join([]string{"DomainIpWhiteListModel", string(data)}, " ")
}
