package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetMissingIndexSwitchNewRequest Request Object
type SetMissingIndexSwitchNewRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *SetMissingIndexSwitchNewRequestBody `json:"body,omitempty"`
}

func (o SetMissingIndexSwitchNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetMissingIndexSwitchNewRequest struct{}"
	}

	return strings.Join([]string{"SetMissingIndexSwitchNewRequest", string(data)}, " ")
}
