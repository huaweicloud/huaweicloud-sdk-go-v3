package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKillProcessTaskRequest Request Object
type ShowKillProcessTaskRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o ShowKillProcessTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKillProcessTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowKillProcessTaskRequest", string(data)}, " ")
}
