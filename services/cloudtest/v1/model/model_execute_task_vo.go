package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteTaskVo 实际的数据类型：单个对象，集合 或 NULL
type ExecuteTaskVo struct {

	// 标志
	Flag *bool `json:"flag,omitempty"`

	// URI
	Uri *string `json:"uri,omitempty"`

	TaskResultVo *TaskResultVo `json:"task_result_vo,omitempty"`

	// 更新用例
	UpdateCaseUriList *[]string `json:"update_case_uri_list,omitempty"`

	// 用例结果列表
	TestCaseResultList *[]TestResultVo `json:"test_case_result_list,omitempty"`
}

func (o ExecuteTaskVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteTaskVo struct{}"
	}

	return strings.Join([]string{"ExecuteTaskVo", string(data)}, " ")
}
