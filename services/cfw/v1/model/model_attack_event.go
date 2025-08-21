package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AttackEvent struct {

	// **参数解释**： 变化数量 **取值范围**： 不涉及
	Changed *int32 `json:"changed,omitempty"`

	// **参数解释**： 阻断数量 **取值范围**： 不涉及
	Deny *int32 `json:"deny,omitempty"`

	// **参数解释**： 总数 **取值范围**： 不涉及
	Total *int32 `json:"total,omitempty"`
}

func (o AttackEvent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttackEvent struct{}"
	}

	return strings.Join([]string{"AttackEvent", string(data)}, " ")
}
