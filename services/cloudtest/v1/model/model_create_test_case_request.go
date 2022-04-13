package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateTestCaseRequest struct {
	Body *CreateTestCaseRequestBody `json:"body,omitempty"`
}

func (o CreateTestCaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTestCaseRequest struct{}"
	}

	return strings.Join([]string{"CreateTestCaseRequest", string(data)}, " ")
}
