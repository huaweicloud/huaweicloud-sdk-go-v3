package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateResolverRuleRouterRequest Request Object
type DisassociateResolverRuleRouterRequest struct {

	// 转发规则ID。
	ResolverruleId string `json:"resolverrule_id"`

	Body *AssociateOrDisassociateRouterWithRuleRequestBody `json:"body,omitempty"`
}

func (o DisassociateResolverRuleRouterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateResolverRuleRouterRequest struct{}"
	}

	return strings.Join([]string{"DisassociateResolverRuleRouterRequest", string(data)}, " ")
}
