package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTestCaseScriptDetailRequest Request Object
type ListTestCaseScriptDetailRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 执行任务id
	TaskId *string `json:"task_id,omitempty"`

	// TMSS用例uri
	TmssCaseUri string `json:"tmss_case_uri"`
}

func (o ListTestCaseScriptDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTestCaseScriptDetailRequest struct{}"
	}

	return strings.Join([]string{"ListTestCaseScriptDetailRequest", string(data)}, " ")
}
