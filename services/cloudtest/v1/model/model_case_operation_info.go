package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CaseOperationInfo DevCloud任务关联用例类
type CaseOperationInfo struct {

	// 用例关联信息
	TestCasesInfo *[]AssignCaseInfo `json:"test_cases_info,omitempty"`

	// 前置用例关联信息
	SetUpCasesInfo *[]AssignCaseInfo `json:"set_up_cases_info,omitempty"`

	// 后置用例关联信息
	TearDownCasesInfo *[]AssignCaseInfo `json:"tear_down_cases_info,omitempty"`
}

func (o CaseOperationInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CaseOperationInfo struct{}"
	}

	return strings.Join([]string{"CaseOperationInfo", string(data)}, " ")
}
