package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResolverRuleRequest Request Object
type ShowResolverRuleRequest struct {

	// 转发规则ID。
	ResolverruleId string `json:"resolverrule_id"`
}

func (o ShowResolverRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResolverRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowResolverRuleRequest", string(data)}, " ")
}
