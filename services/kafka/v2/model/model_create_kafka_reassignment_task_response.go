package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateKafkaReassignmentTaskResponse Response Object
type CreateKafkaReassignmentTaskResponse struct {

	// 任务ID（当执行分区平衡任务时仅返回job_id）。
	JobId *string `json:"job_id,omitempty"`

	// 预估时间，单位为秒（当执行预估时间任务时仅返回reassignment_time）。
	ReassignmentTime *int32 `json:"reassignment_time,omitempty"`
	HttpStatusCode   int    `json:"-"`
}

func (o CreateKafkaReassignmentTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKafkaReassignmentTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateKafkaReassignmentTaskResponse", string(data)}, " ")
}
