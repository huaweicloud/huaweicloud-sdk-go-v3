package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InteractionRuleGroupSummary 互动规则库概要信息
type InteractionRuleGroupSummary struct {

	// 互动规则库ID
	GroupId *string `json:"group_id,omitempty"`

	// 互动规则库名称
	GroupName *string `json:"group_name,omitempty"`

	// 创建时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o InteractionRuleGroupSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InteractionRuleGroupSummary struct{}"
	}

	return strings.Join([]string{"InteractionRuleGroupSummary", string(data)}, " ")
}
