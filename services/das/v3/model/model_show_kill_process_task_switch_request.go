package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKillProcessTaskSwitchRequest Request Object
type ShowKillProcessTaskSwitchRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o ShowKillProcessTaskSwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKillProcessTaskSwitchRequest struct{}"
	}

	return strings.Join([]string{"ShowKillProcessTaskSwitchRequest", string(data)}, " ")
}
