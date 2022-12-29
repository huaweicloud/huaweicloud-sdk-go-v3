package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DisassociateResolveRuleRouterResponse struct {

	// Router(VPC)的ID。
	RouterId *string `json:"router_id,omitempty"`

	// Router(VPC)所在的region。
	RouterRegion *string `json:"router_region,omitempty"`

	// 资源状态。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DisassociateResolveRuleRouterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateResolveRuleRouterResponse struct{}"
	}

	return strings.Join([]string{"DisassociateResolveRuleRouterResponse", string(data)}, " ")
}
