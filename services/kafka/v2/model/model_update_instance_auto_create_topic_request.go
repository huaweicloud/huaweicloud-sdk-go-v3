package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateInstanceAutoCreateTopicRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *UpdateInstanceAutoCreateTopicReq `json:"body,omitempty"`
}

func (o UpdateInstanceAutoCreateTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceAutoCreateTopicRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceAutoCreateTopicRequest", string(data)}, " ")
}
