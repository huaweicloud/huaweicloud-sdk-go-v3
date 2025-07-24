package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetAlarmTopicConfigInfoRequest Request Object
type SetAlarmTopicConfigInfoRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *ConfigAlarmTopicRequest `json:"body,omitempty"`
}

func (o SetAlarmTopicConfigInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAlarmTopicConfigInfoRequest struct{}"
	}

	return strings.Join([]string{"SetAlarmTopicConfigInfoRequest", string(data)}, " ")
}
