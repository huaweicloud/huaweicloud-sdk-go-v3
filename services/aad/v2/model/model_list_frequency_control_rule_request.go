package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFrequencyControlRuleRequest Request Object
type ListFrequencyControlRuleRequest struct {

	// 域名
	DomainName string `json:"domain_name"`
}

func (o ListFrequencyControlRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFrequencyControlRuleRequest struct{}"
	}

	return strings.Join([]string{"ListFrequencyControlRuleRequest", string(data)}, " ")
}
