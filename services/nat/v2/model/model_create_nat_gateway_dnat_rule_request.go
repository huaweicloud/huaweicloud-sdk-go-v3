package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateNatGatewayDnatRuleRequest struct {
	Body *CreateNatGatewayDnatRuleOption `json:"body,omitempty" xml:"body"`
}

func (o CreateNatGatewayDnatRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNatGatewayDnatRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateNatGatewayDnatRuleRequest", string(data)}, " ")
}
