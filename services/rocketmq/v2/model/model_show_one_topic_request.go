package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOneTopicRequest Request Object
type ShowOneTopicRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 主题名称。
	Topic string `json:"topic"`
}

func (o ShowOneTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOneTopicRequest struct{}"
	}

	return strings.Join([]string{"ShowOneTopicRequest", string(data)}, " ")
}
