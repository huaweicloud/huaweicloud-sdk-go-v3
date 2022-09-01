package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteAntitamperRuleResponse struct {

	// 规则id
	Id *string `json:"id,omitempty" xml:"id"`

	// 策略id
	Policyid *string `json:"policyid,omitempty" xml:"policyid"`

	// 防篡改的url
	Url *string `json:"url,omitempty" xml:"url"`

	// 创建规则的时间戳
	Timestamp      *int64 `json:"timestamp,omitempty" xml:"timestamp"`
	HttpStatusCode int    `json:"-"`
}

func (o DeleteAntitamperRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAntitamperRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteAntitamperRuleResponse", string(data)}, " ")
}
