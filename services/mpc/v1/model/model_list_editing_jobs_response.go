package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListEditingJobsResponse struct {

	// 任务列表
	JobArray *[]QueryEditingJob `json:"job_array,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListEditingJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEditingJobsResponse struct{}"
	}

	return strings.Join([]string{"ListEditingJobsResponse", string(data)}, " ")
}
