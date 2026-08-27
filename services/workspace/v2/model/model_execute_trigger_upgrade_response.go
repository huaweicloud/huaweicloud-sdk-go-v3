package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteTriggerUpgradeResponse Response Object
type ExecuteTriggerUpgradeResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ExecuteTriggerUpgradeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteTriggerUpgradeResponse struct{}"
	}

	return strings.Join([]string{"ExecuteTriggerUpgradeResponse", string(data)}, " ")
}
