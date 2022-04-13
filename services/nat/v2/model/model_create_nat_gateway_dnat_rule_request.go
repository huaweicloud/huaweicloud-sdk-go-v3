package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateNatGatewayDnatRuleRequest struct {
	Body *CreateNatGatewayDnatRuleOption `json:"body,omitempty"`
}

func (o CreateNatGatewayDnatRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNatGatewayDnatRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateNatGatewayDnatRuleRequest", string(data)}, " ")
}
