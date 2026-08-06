package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFullDeadLockSwitchNewRequest Request Object
type ShowFullDeadLockSwitchNewRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o ShowFullDeadLockSwitchNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFullDeadLockSwitchNewRequest struct{}"
	}

	return strings.Join([]string{"ShowFullDeadLockSwitchNewRequest", string(data)}, " ")
}
