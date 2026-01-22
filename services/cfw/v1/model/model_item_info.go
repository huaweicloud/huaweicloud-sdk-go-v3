package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ItemInfo struct {

	// **参数解释**： 最大配额 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	MaxQuota *int32 `json:"max_quota,omitempty"`

	// **参数解释**： 已使用配额 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	UsedQuota *int32 `json:"used_quota,omitempty"`

	// **参数解释**： 额外参数，ACL和网络域名使用 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ExtrasInfo map[string]string `json:"extras_info,omitempty"`
}

func (o ItemInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ItemInfo struct{}"
	}

	return strings.Join([]string{"ItemInfo", string(data)}, " ")
}
