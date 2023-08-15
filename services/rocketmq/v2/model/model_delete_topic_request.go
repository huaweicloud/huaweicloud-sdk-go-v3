package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTopicRequest Request Object
type DeleteTopicRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 主题名称。
	Topic string `json:"topic"`
}

func (o DeleteTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTopicRequest struct{}"
	}

	return strings.Join([]string{"DeleteTopicRequest", string(data)}, " ")
}
