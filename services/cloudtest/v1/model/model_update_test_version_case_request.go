package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTestVersionCaseRequest Request Object
type UpdateTestVersionCaseRequest struct {

	// 用例uri
	CaseUri string `json:"case_uri"`

	Body *TestCaseInfo `json:"body,omitempty"`
}

func (o UpdateTestVersionCaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTestVersionCaseRequest struct{}"
	}

	return strings.Join([]string{"UpdateTestVersionCaseRequest", string(data)}, " ")
}
