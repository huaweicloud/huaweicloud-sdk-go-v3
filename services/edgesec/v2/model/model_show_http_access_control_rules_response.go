package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpAccessControlRulesResponse Response Object
type ShowHttpAccessControlRulesResponse struct {

	// 策略下防护规则数量
	Total *int32 `json:"total,omitempty"`

	// 防护规则列表
	Items          *[]ShowHttpAccessControlRuleResponseBody `json:"items,omitempty"`
	HttpStatusCode int                                      `json:"-"`
}

func (o ShowHttpAccessControlRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpAccessControlRulesResponse struct{}"
	}

	return strings.Join([]string{"ShowHttpAccessControlRulesResponse", string(data)}, " ")
}
