package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DisassociateResolveRuleRouterRequest struct {

	// 待解关联resolverrule的ID。
	ResolverruleId string `json:"resolverrule_id"`

	Body *DisassociaterouterReq `json:"body,omitempty"`
}

func (o DisassociateResolveRuleRouterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateResolveRuleRouterRequest struct{}"
	}

	return strings.Join([]string{"DisassociateResolveRuleRouterRequest", string(data)}, " ")
}
