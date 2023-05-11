package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateEditingJobsResponse struct {

	// 接受任务后产生的任务ID。
	JobId *string `json:"job_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateEditingJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEditingJobsResponse struct{}"
	}

	return strings.Join([]string{"CreateEditingJobsResponse", string(data)}, " ")
}
