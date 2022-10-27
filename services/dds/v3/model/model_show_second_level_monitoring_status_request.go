package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowSecondLevelMonitoringStatusRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowSecondLevelMonitoringStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecondLevelMonitoringStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowSecondLevelMonitoringStatusRequest", string(data)}, " ")
}
