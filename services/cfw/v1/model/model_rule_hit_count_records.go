package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RuleHitCountRecords 规则击中次数记录
type RuleHitCountRecords struct {

	// 每页显示个数
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset *int32 `json:"offset,omitempty"`

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 规则击中次数列表
	Records *[]RuleHitCountObject `json:"records,omitempty"`
}

func (o RuleHitCountRecords) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleHitCountRecords struct{}"
	}

	return strings.Join([]string{"RuleHitCountRecords", string(data)}, " ")
}
