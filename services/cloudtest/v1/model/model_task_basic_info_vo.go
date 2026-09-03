package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskBasicInfoVo struct {

	// 任务状态
	ErrorReason *string `json:"error_reason,omitempty"`

	// 任务ID
	Id *string `json:"id,omitempty"`

	// 任务名称
	Name *string `json:"name,omitempty"`

	// 任务类型
	TaskState *int32 `json:"task_state,omitempty"`

	// 测试套类型
	TestSuiteType *int32 `json:"test_suite_type,omitempty"`
}

func (o TaskBasicInfoVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskBasicInfoVo struct{}"
	}

	return strings.Join([]string{"TaskBasicInfoVo", string(data)}, " ")
}
