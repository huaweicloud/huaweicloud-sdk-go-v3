package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInteractionRuleGroupResponse Response Object
type CreateInteractionRuleGroupResponse struct {

	// 互动规则库ID
	GroupId *string `json:"group_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateInteractionRuleGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInteractionRuleGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateInteractionRuleGroupResponse", string(data)}, " ")
}
