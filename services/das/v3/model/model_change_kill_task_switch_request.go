package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeKillTaskSwitchRequest Request Object
type ChangeKillTaskSwitchRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *ChangeKillTaskSwitchRequestBody `json:"body,omitempty"`
}

func (o ChangeKillTaskSwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeKillTaskSwitchRequest struct{}"
	}

	return strings.Join([]string{"ChangeKillTaskSwitchRequest", string(data)}, " ")
}
