package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type CreateCompareTaskResult struct {

	// 对比任务创建成功后，返回对比任务的id，用于查询该对比任务的结果。
	CompareTaskId *string `json:"compare_task_id,omitempty" xml:"compare_task_id"`

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	// 错误信息。
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg"`
}

func (o CreateCompareTaskResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCompareTaskResult struct{}"
	}

	return strings.Join([]string{"CreateCompareTaskResult", string(data)}, " ")
}
