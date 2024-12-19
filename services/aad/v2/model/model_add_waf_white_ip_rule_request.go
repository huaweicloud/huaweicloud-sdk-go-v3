package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddWafWhiteIpRuleRequest Request Object
type AddWafWhiteIpRuleRequest struct {
	Body *AddWafWhiteIpRuleV2RequestBody `json:"body,omitempty"`
}

func (o AddWafWhiteIpRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddWafWhiteIpRuleRequest struct{}"
	}

	return strings.Join([]string{"AddWafWhiteIpRuleRequest", string(data)}, " ")
}
