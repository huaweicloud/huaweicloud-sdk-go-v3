package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHttpBlockTrustIpRuleResponse Response Object
type DeleteHttpBlockTrustIpRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteHttpBlockTrustIpRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHttpBlockTrustIpRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteHttpBlockTrustIpRuleResponse", string(data)}, " ")
}
