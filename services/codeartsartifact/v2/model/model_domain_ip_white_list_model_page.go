package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DomainIpWhiteListModelPage struct {

	// **参数解释**: 总记录数。 **取值范围**: 不涉及。
	TotalRecords *int32 `json:"total_records,omitempty"`

	// **参数解释**: 总页数。 **取值范围**: 不涉及。
	TotalPages *int32 `json:"total_pages,omitempty"`

	// **参数解释**: 租户IP白名单列表。 **取值范围**: 不涉及。
	Data *[]DomainIpWhiteListModel `json:"data,omitempty"`
}

func (o DomainIpWhiteListModelPage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainIpWhiteListModelPage struct{}"
	}

	return strings.Join([]string{"DomainIpWhiteListModelPage", string(data)}, " ")
}
