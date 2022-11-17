package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRuleAclDto
type OrderRuleAclDto struct {

	// 目标规则id，添加规则位于此规则之后，非置顶时不能为空，置顶时为空
	DestRuleId *string `json:"dest_rule_id,omitempty"`

	// 是否置顶，0代表非置顶，1代表置顶
	Top *int32 `json:"top,omitempty"`
}

func (o OrderRuleAclDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrderRuleAclDto struct{}"
	}

	return strings.Join([]string{"OrderRuleAclDto", string(data)}, " ")
}
