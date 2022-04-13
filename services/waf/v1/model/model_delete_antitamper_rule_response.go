package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteAntitamperRuleResponse struct {
	// 规则id

	Id *string `json:"id,omitempty"`
	// 策略id

	Policyid *string `json:"policyid,omitempty"`
	// 防篡改的url

	Url *string `json:"url,omitempty"`
	// 创建规则的时间戳

	Timestamp      *int64 `json:"timestamp,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o DeleteAntitamperRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAntitamperRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteAntitamperRuleResponse", string(data)}, " ")
}
