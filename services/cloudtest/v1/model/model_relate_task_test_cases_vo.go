package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RelateTaskTestCasesVo 用例及任务信息
type RelateTaskTestCasesVo struct {

	// 用例编号
	TestCaseNum *string `json:"test_case_num,omitempty"`

	// 用例名
	TestCaseName *string `json:"test_case_name,omitempty"`

	// 用例uri
	TestCaseUri *string `json:"test_case_uri,omitempty"`

	// 任务uri
	TaskUri *string `json:"task_uri,omitempty"`

	// 任务名
	TaskName *string `json:"task_name,omitempty"`

	// 任务编号
	TaskNum *string `json:"task_num,omitempty"`

	// 任务创建人
	TaskCreator *string `json:"task_creator,omitempty"`
}

func (o RelateTaskTestCasesVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RelateTaskTestCasesVo struct{}"
	}

	return strings.Join([]string{"RelateTaskTestCasesVo", string(data)}, " ")
}
