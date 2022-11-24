package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateResolveRuleRequest struct {
	Body *ResolveRuleReq `json:"body,omitempty"`
}

func (o CreateResolveRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResolveRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateResolveRuleRequest", string(data)}, " ")
}
