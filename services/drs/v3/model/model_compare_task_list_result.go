package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type CompareTaskListResult struct {

	// 对比任务列表。
	CompareTaskList *[]CompareTaskList `json:"CompareTaskList,omitempty" xml:"CompareTaskList"`

	// 对比任务列表总数。
	CompareTaskListCount *int32 `json:"compare_task_list_count,omitempty" xml:"compare_task_list_count"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg"`

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`
}

func (o CompareTaskListResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareTaskListResult struct{}"
	}

	return strings.Join([]string{"CompareTaskListResult", string(data)}, " ")
}
