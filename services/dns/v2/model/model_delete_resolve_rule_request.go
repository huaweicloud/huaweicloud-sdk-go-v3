package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteResolveRuleRequest struct {

	// 待删除resolverrule的ID。
	ResolverruleId string `json:"resolverrule_id"`
}

func (o DeleteResolveRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResolveRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteResolveRuleRequest", string(data)}, " ")
}
