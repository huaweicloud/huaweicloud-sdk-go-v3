package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRuleRequest Request Object
type UpdateRuleRequest struct {

	// 权限策略id
	Ruleid string `json:"ruleid"`

	Body *UpdateRuleRequestBody `json:"body,omitempty"`
}

func (o UpdateRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateRuleRequest", string(data)}, " ")
}
