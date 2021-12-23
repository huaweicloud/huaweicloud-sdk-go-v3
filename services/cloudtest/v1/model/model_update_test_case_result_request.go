package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateTestCaseResultRequest struct {
	Body *UpdateTestCaseResultRequestBody `json:"body,omitempty"`
}

func (o UpdateTestCaseResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTestCaseResultRequest struct{}"
	}

	return strings.Join([]string{"UpdateTestCaseResultRequest", string(data)}, " ")
}
