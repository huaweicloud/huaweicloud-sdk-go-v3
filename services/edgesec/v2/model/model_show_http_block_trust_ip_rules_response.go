package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpBlockTrustIpRulesResponse Response Object
type ShowHttpBlockTrustIpRulesResponse struct {

	// 策略下防护规则数量
	Total *int32 `json:"total,omitempty"`

	// 防护规则列表
	Items *[]ShowHttpBlockTrustIpRuleResponseBody `json:"items,omitempty"`

	// 防护规则下IP/IP段数量
	Size           *int64 `json:"size,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowHttpBlockTrustIpRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpBlockTrustIpRulesResponse struct{}"
	}

	return strings.Join([]string{"ShowHttpBlockTrustIpRulesResponse", string(data)}, " ")
}
