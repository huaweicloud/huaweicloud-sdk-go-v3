package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AssociateResolveRuleRouterRequest struct {

	// 关联VPC的resolverrule的ID。
	ResolverruleId string `json:"resolverrule_id"`

	Body *AssociateRouterReq `json:"body,omitempty"`
}

func (o AssociateResolveRuleRouterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateResolveRuleRouterRequest struct{}"
	}

	return strings.Join([]string{"AssociateResolveRuleRouterRequest", string(data)}, " ")
}
