package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWafWhiteIpRuleRequest Request Object
type DeleteWafWhiteIpRuleRequest struct {
	Body *DeleteWafWhiteIpRuleV2RequestBody `json:"body,omitempty"`
}

func (o DeleteWafWhiteIpRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWafWhiteIpRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteWafWhiteIpRuleRequest", string(data)}, " ")
}
