package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowConfigQuotaDto struct {
	ItemInfo *ItemInfo `json:"item_info,omitempty"`

	// **参数解释**： 防火墙成员最大配额 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	MaxQuota *int32 `json:"max_quota,omitempty"`

	// **参数解释**： 防火墙成员配额类型 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	QuotaType *string `json:"quota_type,omitempty"`

	// **参数解释**： 已使用配额 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	UsedQuota *int32 `json:"used_quota,omitempty"`
}

func (o ShowConfigQuotaDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigQuotaDto struct{}"
	}

	return strings.Join([]string{"ShowConfigQuotaDto", string(data)}, " ")
}
