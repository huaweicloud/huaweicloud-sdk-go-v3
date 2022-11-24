package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateNatGatewaySnatRuleRequest struct {
	Body *CreateNatGatewaySnatRuleRequestOption `json:"body,omitempty"`
}

func (o CreateNatGatewaySnatRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNatGatewaySnatRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateNatGatewaySnatRuleRequest", string(data)}, " ")
}
