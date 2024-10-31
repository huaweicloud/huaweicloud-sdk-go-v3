package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpCcRulesResponse Response Object
type ShowHttpCcRulesResponse struct {

	// 策略下cc规则数量
	Total *int32 `json:"total,omitempty"`

	// cc规则列表
	Items          *[]ShowHttpCcRuleResponseBody `json:"items,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ShowHttpCcRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpCcRulesResponse struct{}"
	}

	return strings.Join([]string{"ShowHttpCcRulesResponse", string(data)}, " ")
}
