package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkJobList 作业列表详情。
type FlinkJobList struct {

	// 作业查询结果条数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 作业信息
	Jobs *[]FlinkJobInfo `json:"jobs,omitempty"`
}

func (o FlinkJobList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkJobList struct{}"
	}

	return strings.Join([]string{"FlinkJobList", string(data)}, " ")
}
