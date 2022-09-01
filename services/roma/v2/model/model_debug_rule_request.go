package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DebugRuleRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *DebugRuleRequestBody `json:"body,omitempty" xml:"body"`
}

func (o DebugRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DebugRuleRequest struct{}"
	}

	return strings.Join([]string{"DebugRuleRequest", string(data)}, " ")
}
