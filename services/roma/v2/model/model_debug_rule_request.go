package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DebugRuleRequest Request Object
type DebugRuleRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *DebugRuleRequestBody `json:"body,omitempty"`
}

func (o DebugRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DebugRuleRequest struct{}"
	}

	return strings.Join([]string{"DebugRuleRequest", string(data)}, " ")
}
