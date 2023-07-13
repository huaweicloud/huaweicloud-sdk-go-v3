package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunTestCaseRequest Request Object
type RunTestCaseRequest struct {
	Body *RunTestCaseRequestBody `json:"body,omitempty"`
}

func (o RunTestCaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunTestCaseRequest struct{}"
	}

	return strings.Join([]string{"RunTestCaseRequest", string(data)}, " ")
}
