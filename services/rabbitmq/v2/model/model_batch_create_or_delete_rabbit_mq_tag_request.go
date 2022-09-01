package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchCreateOrDeleteRabbitMqTagRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *BatchCreateOrDeleteTagReq `json:"body,omitempty" xml:"body"`
}

func (o BatchCreateOrDeleteRabbitMqTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateOrDeleteRabbitMqTagRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateOrDeleteRabbitMqTagRequest", string(data)}, " ")
}
