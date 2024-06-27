package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTestCaseCommentResponse Response Object
type DeleteTestCaseCommentResponse struct {

	// 对外时：success|error;
	Status *string `json:"status,omitempty"`

	Result         *ResultValueStringForOk `json:"result,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o DeleteTestCaseCommentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTestCaseCommentResponse struct{}"
	}

	return strings.Join([]string{"DeleteTestCaseCommentResponse", string(data)}, " ")
}
