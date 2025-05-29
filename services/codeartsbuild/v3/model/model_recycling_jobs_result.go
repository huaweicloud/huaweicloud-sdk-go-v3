package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecyclingJobsResult 结果
type RecyclingJobsResult struct {

	// 任务保留时间
	KeepTime *int32 `json:"keep_time,omitempty"`

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 任务列表
	JobList *[]RecyclingJob `json:"job_list,omitempty"`
}

func (o RecyclingJobsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecyclingJobsResult struct{}"
	}

	return strings.Join([]string{"RecyclingJobsResult", string(data)}, " ")
}
