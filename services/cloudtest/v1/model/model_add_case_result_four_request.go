package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddCaseResultFourRequest Request Object
type AddCaseResultFourRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 版本ID
	VersionUri string `json:"version_uri"`

	// 用例uri
	CaseUri string `json:"case_uri"`

	Body *AddTestCaseResultInfo `json:"body,omitempty"`
}

func (o AddCaseResultFourRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddCaseResultFourRequest struct{}"
	}

	return strings.Join([]string{"AddCaseResultFourRequest", string(data)}, " ")
}
