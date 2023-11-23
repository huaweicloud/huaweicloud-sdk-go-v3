package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKillOpRuleRuleListResponse Response Object
type ShowKillOpRuleRuleListResponse struct {

	// 规则列表。
	Rules *[]KillOpRule `json:"rules,omitempty"`

	// 总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowKillOpRuleRuleListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKillOpRuleRuleListResponse struct{}"
	}

	return strings.Join([]string{"ShowKillOpRuleRuleListResponse", string(data)}, " ")
}
