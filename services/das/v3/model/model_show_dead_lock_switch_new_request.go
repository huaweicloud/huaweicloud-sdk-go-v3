package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeadLockSwitchNewRequest Request Object
type ShowDeadLockSwitchNewRequest struct {

	// 引擎类型
	EngineType *string `json:"engine_type,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`
}

func (o ShowDeadLockSwitchNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeadLockSwitchNewRequest struct{}"
	}

	return strings.Join([]string{"ShowDeadLockSwitchNewRequest", string(data)}, " ")
}
