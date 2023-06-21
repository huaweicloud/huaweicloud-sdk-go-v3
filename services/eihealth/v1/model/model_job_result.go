package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 作业运行结果信息
type JobResult struct {

	// 输入总数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 失败个数
	FailedCount *int32 `json:"failed_count,omitempty"`
}

func (o JobResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobResult struct{}"
	}

	return strings.Join([]string{"JobResult", string(data)}, " ")
}
