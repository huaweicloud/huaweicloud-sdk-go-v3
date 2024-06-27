package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskAssignCaseVo 实际的数据类型：单个对象，集合 或 NULL
type TaskAssignCaseVo struct {

	// 排序顺序
	Sort *int32 `json:"sort,omitempty"`

	// 用例uri
	CaseUri *string `json:"case_uri,omitempty"`

	// 是否可用
	IsAvailable *int32 `json:"is_available,omitempty"`

	// 用例名称
	TestCaseName *string `json:"test_case_name,omitempty"`

	// 用例编号
	TestCaseNumber *string `json:"test_case_number,omitempty"`
}

func (o TaskAssignCaseVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskAssignCaseVo struct{}"
	}

	return strings.Join([]string{"TaskAssignCaseVo", string(data)}, " ")
}
