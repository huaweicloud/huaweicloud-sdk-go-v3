package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetFullDeadLockSwitchNewRequest Request Object
type SetFullDeadLockSwitchNewRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *SetFullDeadLockSwitchNewRequestBody `json:"body,omitempty"`
}

func (o SetFullDeadLockSwitchNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetFullDeadLockSwitchNewRequest struct{}"
	}

	return strings.Join([]string{"SetFullDeadLockSwitchNewRequest", string(data)}, " ")
}
