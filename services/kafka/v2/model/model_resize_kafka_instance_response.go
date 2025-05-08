package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeKafkaInstanceResponse Response Object
type ResizeKafkaInstanceResponse struct {

	// 规格变更任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResizeKafkaInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeKafkaInstanceResponse struct{}"
	}

	return strings.Join([]string{"ResizeKafkaInstanceResponse", string(data)}, " ")
}
