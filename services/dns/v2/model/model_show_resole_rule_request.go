package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowResoleRuleRequest struct {

	// 待查询resolverrule的ID。
	ResolverruleId string `json:"resolverrule_id"`
}

func (o ShowResoleRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResoleRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowResoleRuleRequest", string(data)}, " ")
}
