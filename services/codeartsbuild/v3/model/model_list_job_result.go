package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobResult 返回结果
type ListJobResult struct {

	// 任务总数
	Total *int32 `json:"total,omitempty"`

	// 任务列表
	JobList *[]ListJobResultJobList `json:"job_list,omitempty"`
}

func (o ListJobResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobResult struct{}"
	}

	return strings.Join([]string{"ListJobResult", string(data)}, " ")
}
