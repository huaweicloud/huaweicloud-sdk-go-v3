package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobs2Response Response Object
type ListJobs2Response struct {

	// 任务总数。
	JobCount *int32 `json:"job_count,omitempty"`

	// 任务列表。
	JobList        *[]ListJobsRespJobList `json:"job_list,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListJobs2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobs2Response struct{}"
	}

	return strings.Join([]string{"ListJobs2Response", string(data)}, " ")
}
