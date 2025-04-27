package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateResolverRuleRequest Request Object
type UpdateResolverRuleRequest struct {

	// 转发规则ID。
	ResolverruleId string `json:"resolverrule_id"`

	Body *UpdateResolverRuleRequestBody `json:"body,omitempty"`
}

func (o UpdateResolverRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResolverRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateResolverRuleRequest", string(data)}, " ")
}
