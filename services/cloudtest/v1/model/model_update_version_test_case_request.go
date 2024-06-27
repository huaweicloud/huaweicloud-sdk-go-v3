package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVersionTestCaseRequest Request Object
type UpdateVersionTestCaseRequest struct {

	// 用例uri
	CaseId string `json:"case_id"`

	Body *TestCaseInfo `json:"body,omitempty"`
}

func (o UpdateVersionTestCaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVersionTestCaseRequest struct{}"
	}

	return strings.Join([]string{"UpdateVersionTestCaseRequest", string(data)}, " ")
}
