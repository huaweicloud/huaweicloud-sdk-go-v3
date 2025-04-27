package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResolverRuleRequest Request Object
type CreateResolverRuleRequest struct {
	Body *CreateResolverRuleRequestBody `json:"body,omitempty"`
}

func (o CreateResolverRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResolverRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateResolverRuleRequest", string(data)}, " ")
}
