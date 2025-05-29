package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskResultDetailVo struct {

	// 处理人名称
	Owner *string `json:"owner,omitempty"`

	// 用例结果统计信息
	CaseResultStatics *interface{} `json:"case_result_statics,omitempty"`

	TaskResult *TaskResultVo `json:"task_result,omitempty"`

	// 实际的数据类型：单个对象，集合 或 NULL
	TestResultList *[]TaskResultVo `json:"test_result_list,omitempty"`
}

func (o TaskResultDetailVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskResultDetailVo struct{}"
	}

	return strings.Join([]string{"TaskResultDetailVo", string(data)}, " ")
}
