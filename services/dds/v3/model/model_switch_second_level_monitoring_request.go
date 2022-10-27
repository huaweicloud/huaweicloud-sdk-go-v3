package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SwitchSecondLevelMonitoringRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *SwitchSecondLevelMonitoringRequestBody `json:"body,omitempty"`
}

func (o SwitchSecondLevelMonitoringRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchSecondLevelMonitoringRequest struct{}"
	}

	return strings.Join([]string{"SwitchSecondLevelMonitoringRequest", string(data)}, " ")
}
