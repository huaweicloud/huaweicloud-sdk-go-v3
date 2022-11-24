package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateStreamingJobResponse struct {

	// 作业ID
	JobId *string `json:"job_id,omitempty"`

	// 作业状态
	JobState *string `json:"job_state,omitempty"`

	// 操作结果
	Status *string `json:"status,omitempty"`

	// 作业错误详情
	CheckInfo      map[string]interface{} `json:"check_info,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o UpdateStreamingJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStreamingJobResponse struct{}"
	}

	return strings.Join([]string{"UpdateStreamingJobResponse", string(data)}, " ")
}
