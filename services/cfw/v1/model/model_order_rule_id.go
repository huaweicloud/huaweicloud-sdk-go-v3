package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OrderRuleId struct {

	// **参数解释**： 移动的规则ID **取值范围**： 不涉及
	Id *string `json:"id,omitempty"`
}

func (o OrderRuleId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrderRuleId struct{}"
	}

	return strings.Join([]string{"OrderRuleId", string(data)}, " ")
}
