package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// entities字段数据结构说明
type Entities struct {

	// 子任务数量。没有子任务时为0
	SubJobsTotal *int32 `json:"sub_jobs_total,omitempty" xml:"sub_jobs_total"`

	// 每个子任务的执行信息。没有子任务时为空列表
	SubJobs *[]SubJobs `json:"sub_jobs,omitempty" xml:"sub_jobs"`
}

func (o Entities) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Entities struct{}"
	}

	return strings.Join([]string{"Entities", string(data)}, " ")
}
