package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListFailureJobsResponse struct {

	// 失败任务信息列表。
	FailureJobs *[]FailureJobParams `json:"failure_jobs,omitempty" xml:"failure_jobs"`

	// 列表中失败任务个数。
	Count          *int32 `json:"count,omitempty" xml:"count"`
	HttpStatusCode int    `json:"-"`
}

func (o ListFailureJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFailureJobsResponse struct{}"
	}

	return strings.Join([]string{"ListFailureJobsResponse", string(data)}, " ")
}
