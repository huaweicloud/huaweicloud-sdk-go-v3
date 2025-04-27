package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateResolverRuleRouterResponse Response Object
type AssociateResolverRuleRouterResponse struct {

	// 关联VPC的ID。
	RouterId *string `json:"router_id,omitempty"`

	// 关联VPC所在的region。
	RouterRegion *string `json:"router_region,omitempty"`

	// 资源状态。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AssociateResolverRuleRouterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateResolverRuleRouterResponse struct{}"
	}

	return strings.Join([]string{"AssociateResolverRuleRouterResponse", string(data)}, " ")
}
