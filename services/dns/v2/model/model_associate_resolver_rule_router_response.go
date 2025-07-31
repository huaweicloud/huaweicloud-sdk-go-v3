package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateResolverRuleRouterResponse Response Object
type AssociateResolverRuleRouterResponse struct {

	// **参数解释：** 关联VPC的ID。 **取值范围：** 不涉及。
	RouterId *string `json:"router_id,omitempty"`

	// **参数解释：** 关联VPC所在的region。 **取值范围：** 不涉及。
	RouterRegion *string `json:"router_region,omitempty"`

	// **参数解释：** 关联VPC的状态。 **取值范围：** 不涉及。
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
