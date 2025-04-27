package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateResolverRuleRouterRequest Request Object
type AssociateResolverRuleRouterRequest struct {

	// 转发规则ID。
	ResolverruleId string `json:"resolverrule_id"`

	Body *AssociateOrDisassociateRouterWithRuleRequestBody `json:"body,omitempty"`
}

func (o AssociateResolverRuleRouterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateResolverRuleRouterRequest struct{}"
	}

	return strings.Join([]string{"AssociateResolverRuleRouterRequest", string(data)}, " ")
}
