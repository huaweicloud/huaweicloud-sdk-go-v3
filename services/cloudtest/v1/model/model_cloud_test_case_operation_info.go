package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CloudTestCaseOperationInfo struct {

	// 前置用例信息
	SetUpCasesInfo *[]CloudTestCaseInfo `json:"setUpCasesInfo,omitempty"`

	// 后置用例信息
	TearDownCasesInfo *[]CloudTestCaseInfo `json:"tearDownCasesInfo,omitempty"`

	// 用例信息
	TestCasesInfo *[]CloudTestCaseInfo `json:"testCasesInfo,omitempty"`
}

func (o CloudTestCaseOperationInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudTestCaseOperationInfo struct{}"
	}

	return strings.Join([]string{"CloudTestCaseOperationInfo", string(data)}, " ")
}
