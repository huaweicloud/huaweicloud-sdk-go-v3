package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateResolveRuleRequest struct {

	// 待修改resolverrule的ID。
	ResolverruleId string `json:"resolverrule_id"`

	Body *UpdateResolveRuleReq `json:"body,omitempty"`
}

func (o UpdateResolveRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResolveRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateResolveRuleRequest", string(data)}, " ")
}
