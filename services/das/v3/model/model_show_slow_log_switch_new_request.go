package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSlowLogSwitchNewRequest Request Object
type ShowSlowLogSwitchNewRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 引擎类型
	EngineType *string `json:"engine_type,omitempty"`
}

func (o ShowSlowLogSwitchNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSlowLogSwitchNewRequest struct{}"
	}

	return strings.Join([]string{"ShowSlowLogSwitchNewRequest", string(data)}, " ")
}
