package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OpensourceCount struct {

	// **参数解释**: 危急漏洞数。 **取值范围**: 不涉及。
	Critical *int32 `json:"critical,omitempty"`

	// **参数解释**: 高危漏洞数。 **取值范围**: 不涉及。
	High *int32 `json:"high,omitempty"`

	// **参数解释**: 中危漏洞数。 **取值范围**: 不涉及。
	Medium *int32 `json:"medium,omitempty"`

	// **参数解释**: 低危漏洞数。 **取值范围**: 不涉及。
	Low *int32 `json:"low,omitempty"`
}

func (o OpensourceCount) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpensourceCount struct{}"
	}

	return strings.Join([]string{"OpensourceCount", string(data)}, " ")
}
