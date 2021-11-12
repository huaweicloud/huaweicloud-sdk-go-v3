package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 调用成功时表示调用结果。  调用失败时无此字段。
type TaskSumbitResponseResult struct {
	// 批量图像内容审核的任务标识，用于后续的结果查询。

	JobId *string `json:"job_id,omitempty"`
}

func (o TaskSumbitResponseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskSumbitResponseResult struct{}"
	}

	return strings.Join([]string{"TaskSumbitResponseResult", string(data)}, " ")
}
