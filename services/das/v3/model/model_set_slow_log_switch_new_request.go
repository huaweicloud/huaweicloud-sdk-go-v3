package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetSlowLogSwitchNewRequest Request Object
type SetSlowLogSwitchNewRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *SetSlowLogSwitchNewRequestBody `json:"body,omitempty"`
}

func (o SetSlowLogSwitchNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetSlowLogSwitchNewRequest struct{}"
	}

	return strings.Join([]string{"SetSlowLogSwitchNewRequest", string(data)}, " ")
}
