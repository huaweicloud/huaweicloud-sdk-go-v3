package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CaseOperationVo 任务关联用例信息
type CaseOperationVo struct {

	// 用例关联信息
	TestCasesInfo *[]TaskAssignCaseVo `json:"test_cases_info,omitempty"`

	// 前置用例关联信息
	SetUpCasesInfo *[]TaskAssignCaseVo `json:"set_up_cases_info,omitempty"`

	// 后置用例关联信息
	TearDownCasesInfo *[]TaskAssignCaseVo `json:"tear_down_cases_info,omitempty"`
}

func (o CaseOperationVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CaseOperationVo struct{}"
	}

	return strings.Join([]string{"CaseOperationVo", string(data)}, " ")
}
