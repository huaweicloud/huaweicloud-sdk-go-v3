package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWafCustomRuleRequest Request Object
type ListWafCustomRuleRequest struct {

	// 域名
	DomainName string `json:"domain_name"`

	// 防护区域，0-大陆，1-海外
	OverseasType int32 `json:"overseas_type"`
}

func (o ListWafCustomRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWafCustomRuleRequest struct{}"
	}

	return strings.Join([]string{"ListWafCustomRuleRequest", string(data)}, " ")
}
