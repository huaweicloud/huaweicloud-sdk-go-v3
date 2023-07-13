package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateConsumerGroupResponse Response Object
type BatchUpdateConsumerGroupResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchUpdateConsumerGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateConsumerGroupResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateConsumerGroupResponse", string(data)}, " ")
}
