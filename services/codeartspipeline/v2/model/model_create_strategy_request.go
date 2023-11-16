package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStrategyRequest Request Object
type CreateStrategyRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	Body *CreateRuleSetReq `json:"body,omitempty"`
}

func (o CreateStrategyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStrategyRequest struct{}"
	}

	return strings.Join([]string{"CreateStrategyRequest", string(data)}, " ")
}
