package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowRabbitMqTagsRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
}

func (o ShowRabbitMqTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRabbitMqTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowRabbitMqTagsRequest", string(data)}, " ")
}
