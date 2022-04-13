package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type CompareTaskListResult struct {
	// 对比任务列表。

	CompareTaskList *[]CompareTaskList `json:"CompareTaskList,omitempty"`
	// 对比任务列表总数。

	CompareTaskListCount *int32 `json:"compare_task_list_count,omitempty"`
	// 错误信息

	ErrorMsg *string `json:"error_msg,omitempty"`
	// 错误码。

	ErrorCode *string `json:"error_code,omitempty"`
}

func (o CompareTaskListResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareTaskListResult struct{}"
	}

	return strings.Join([]string{"CompareTaskListResult", string(data)}, " ")
}
