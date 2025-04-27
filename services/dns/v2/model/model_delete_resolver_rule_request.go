package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteResolverRuleRequest Request Object
type DeleteResolverRuleRequest struct {

	// 转发规则ID。
	ResolverruleId string `json:"resolverrule_id"`
}

func (o DeleteResolverRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResolverRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteResolverRuleRequest", string(data)}, " ")
}
