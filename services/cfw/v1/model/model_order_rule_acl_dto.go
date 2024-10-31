package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OrderRuleAclDto UpdateRuleAclDto
type OrderRuleAclDto struct {

	// 目标规则id，添加规则位于此规则之后，非置顶时不能为空，置顶时为空，目标规则id可以通过[查询防护规则接口](ListAclRules.xml)获得，通过返回值中的data.records.rule_id（.表示各对象之间层级的区分）获得。
	DestRuleId *string `json:"dest_rule_id,omitempty"`

	// 是否置顶，0代表非置顶，1代表置顶
	Top *int32 `json:"top,omitempty"`

	// 是否置底，0代表非置底，1代表置底
	Bottom *int32 `json:"bottom,omitempty"`
}

func (o OrderRuleAclDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrderRuleAclDto struct{}"
	}

	return strings.Join([]string{"OrderRuleAclDto", string(data)}, " ")
}
