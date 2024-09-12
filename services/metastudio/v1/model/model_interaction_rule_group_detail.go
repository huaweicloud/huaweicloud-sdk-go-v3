package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InteractionRuleGroupDetail 互动规则库信息
type InteractionRuleGroupDetail struct {

	// 互动规则库ID
	GroupId *string `json:"group_id,omitempty"`

	// 互动规则库名称
	GroupName string `json:"group_name"`

	// 互动规则列表
	InteractionRules *[]InteractionRuleDetailInfo `json:"interaction_rules,omitempty"`

	// 创建时间，格式遵循：RFC 3339 如“2021-01-10T08:43:17Z”。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间，格式遵循：RFC 3339 如“2021-01-10T08:43:17Z”。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o InteractionRuleGroupDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InteractionRuleGroupDetail struct{}"
	}

	return strings.Join([]string{"InteractionRuleGroupDetail", string(data)}, " ")
}
