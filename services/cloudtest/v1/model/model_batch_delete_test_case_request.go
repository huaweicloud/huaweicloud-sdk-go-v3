package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchDeleteTestCaseRequest struct {
	Body *BatchDeleteTestCaseRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteTestCaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteTestCaseRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteTestCaseRequest", string(data)}, " ")
}
