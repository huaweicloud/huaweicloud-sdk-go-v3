package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateConsumerGroupRequest struct {

	// 指定的队列ID
	QueueId string `json:"queue_id" xml:"queue_id"`

	Body *CreateConsumerGroupReq `json:"body,omitempty" xml:"body"`
}

func (o CreateConsumerGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConsumerGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateConsumerGroupRequest", string(data)}, " ")
}
