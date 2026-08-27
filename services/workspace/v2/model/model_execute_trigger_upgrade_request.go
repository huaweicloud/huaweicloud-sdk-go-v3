package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteTriggerUpgradeRequest Request Object
type ExecuteTriggerUpgradeRequest struct {
	Body *ExecuteTriggerUpgradeRequestBody `json:"body,omitempty"`
}

func (o ExecuteTriggerUpgradeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteTriggerUpgradeRequest struct{}"
	}

	return strings.Join([]string{"ExecuteTriggerUpgradeRequest", string(data)}, " ")
}
