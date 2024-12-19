package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddWafWhiteIpRuleResponse Response Object
type AddWafWhiteIpRuleResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddWafWhiteIpRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddWafWhiteIpRuleResponse struct{}"
	}

	return strings.Join([]string{"AddWafWhiteIpRuleResponse", string(data)}, " ")
}
