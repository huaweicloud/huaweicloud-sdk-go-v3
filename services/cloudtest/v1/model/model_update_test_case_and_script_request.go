package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTestCaseAndScriptRequest Request Object
type UpdateTestCaseAndScriptRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// TMSS用例uri
	TmssCaseUri string `json:"tmss_case_uri"`

	// 新组合AW开关
	TurnOnAwmapping *bool `json:"turn_on_awmapping,omitempty"`

	Body *CreateTestCaseReq `json:"body,omitempty"`
}

func (o UpdateTestCaseAndScriptRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTestCaseAndScriptRequest struct{}"
	}

	return strings.Join([]string{"UpdateTestCaseAndScriptRequest", string(data)}, " ")
}
