package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type JobEntities struct {

	// 每个子任务的执行信息。
	SubJobs *[]SubJob `json:"sub_jobs,omitempty" xml:"sub_jobs"`

	// 子任务数量。
	SubJobsTotal *int32 `json:"sub_jobs_total,omitempty" xml:"sub_jobs_total"`
}

func (o JobEntities) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobEntities struct{}"
	}

	return strings.Join([]string{"JobEntities", string(data)}, " ")
}
