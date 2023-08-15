package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTopicRequest Request Object
type UpdateTopicRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 主题名称。
	Topic string `json:"topic"`

	Body *UpdateTopicReq `json:"body,omitempty"`
}

func (o UpdateTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTopicRequest struct{}"
	}

	return strings.Join([]string{"UpdateTopicRequest", string(data)}, " ")
}
