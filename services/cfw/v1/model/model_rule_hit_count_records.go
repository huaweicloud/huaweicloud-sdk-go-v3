package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RuleHitCountRecords **参数解释**： 规则击中次数记录
type RuleHitCountRecords struct {

	// **参数解释**： 每页显示个数 **取值范围**： 1-1024
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 偏移量：指定返回记录的开始位置，必须为数字 **取值范围**： 大于或等于0，默认0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 获取规则击中次数总条数 **取值范围**： 不涉及
	Total *int32 `json:"total,omitempty"`

	// **参数解释**： 规则击中次数信息列表 **约束限制**： 不涉及
	Records *[]RuleHitCountObject `json:"records,omitempty"`
}

func (o RuleHitCountRecords) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleHitCountRecords struct{}"
	}

	return strings.Join([]string{"RuleHitCountRecords", string(data)}, " ")
}
