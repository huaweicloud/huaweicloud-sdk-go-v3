package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpIgnoreRulesResponse Response Object
type ShowHttpIgnoreRulesResponse struct {

	// 策略下防护规则数量
	Total *int32 `json:"total,omitempty"`

	// 防护规则列表
	Items          *[]ShowHttpIgnoreRuleResponseBody `json:"items,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ShowHttpIgnoreRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpIgnoreRulesResponse struct{}"
	}

	return strings.Join([]string{"ShowHttpIgnoreRulesResponse", string(data)}, " ")
}
