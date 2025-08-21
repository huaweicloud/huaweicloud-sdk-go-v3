package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangedVo struct {

	// **参数解释**： 变化数量 **取值范围**： 不涉及
	Changed *int32 `json:"changed,omitempty"`

	// **参数解释**： 变化数量 **取值范围**： 不涉及
	Total *int32 `json:"total,omitempty"`

	// **参数解释**： 数量 **取值范围**： 不涉及
	Value *int32 `json:"value,omitempty"`
}

func (o ChangedVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangedVo struct{}"
	}

	return strings.Join([]string{"ChangedVo", string(data)}, " ")
}
