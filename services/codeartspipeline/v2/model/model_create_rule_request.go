package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRuleRequest Request Object
type CreateRuleRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	Body *CreateRuleReq `json:"body,omitempty"`
}

func (o CreateRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateRuleRequest", string(data)}, " ")
}
