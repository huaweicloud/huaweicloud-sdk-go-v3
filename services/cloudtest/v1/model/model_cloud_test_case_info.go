package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CloudTestCaseInfo struct {

	// 用例id
	CaseId *string `json:"case_id,omitempty"`

	// tmss用例类型
	CaseType *int32 `json:"caseType,omitempty"`

	// 是否未禁用，1为未禁用，0为已禁用
	IsForbidden *int32 `json:"is_forbidden,omitempty"`

	Owner *CommonDto `json:"owner,omitempty"`

	Result *CommonDto `json:"result,omitempty"`

	// 用例脚本路径
	ScriptUrl *string `json:"scriptUrl,omitempty"`

	Status *CommonDto `json:"status,omitempty"`

	// 用例名称
	TestCaseName *string `json:"testCaseName,omitempty"`

	// 用例编号
	TestCaseNumber *string `json:"testCaseNumber,omitempty"`

	// tmss版本地址
	TmssVersionUri *string `json:"tmssVersionUri,omitempty"`
}

func (o CloudTestCaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudTestCaseInfo struct{}"
	}

	return strings.Join([]string{"CloudTestCaseInfo", string(data)}, " ")
}
