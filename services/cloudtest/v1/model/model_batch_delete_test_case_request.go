package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchDeleteTestCaseRequest struct {
	Body *BatchDeleteTestCaseRequestBody `json:"body,omitempty" xml:"body"`
}

func (o BatchDeleteTestCaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteTestCaseRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteTestCaseRequest", string(data)}, " ")
}
