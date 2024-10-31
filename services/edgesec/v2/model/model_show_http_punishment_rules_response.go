package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpPunishmentRulesResponse Response Object
type ShowHttpPunishmentRulesResponse struct {

	// 策略下防护规则数量
	Total *int32 `json:"total,omitempty"`

	// 防护规则列表
	Items          *[]ShowHttpPunishmentRuleResponseBody `json:"items,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o ShowHttpPunishmentRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpPunishmentRulesResponse struct{}"
	}

	return strings.Join([]string{"ShowHttpPunishmentRulesResponse", string(data)}, " ")
}
