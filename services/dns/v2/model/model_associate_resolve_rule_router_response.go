package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AssociateResolveRuleRouterResponse struct {

	// Router(VPC)的ID。
	RouterId *string `json:"router_id,omitempty"`

	// Router(VPC)所在的region。
	RouterRegion *string `json:"router_region,omitempty"`

	// 资源状态。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AssociateResolveRuleRouterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateResolveRuleRouterResponse struct{}"
	}

	return strings.Join([]string{"AssociateResolveRuleRouterResponse", string(data)}, " ")
}
