package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWafWhiteIpRuleResponse Response Object
type DeleteWafWhiteIpRuleResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteWafWhiteIpRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWafWhiteIpRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteWafWhiteIpRuleResponse", string(data)}, " ")
}
