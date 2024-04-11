package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataLevelTableCompareJobsResponse Response Object
type ListDataLevelTableCompareJobsResponse struct {

	// 表对比任务信息
	CompareJobs *[]CompareJobInfo `json:"compare_jobs,omitempty"`

	// 任务数量
	Count          *int64 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDataLevelTableCompareJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataLevelTableCompareJobsResponse struct{}"
	}

	return strings.Join([]string{"ListDataLevelTableCompareJobsResponse", string(data)}, " ")
}
