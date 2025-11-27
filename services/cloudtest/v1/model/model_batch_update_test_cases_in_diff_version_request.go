package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateTestCasesInDiffVersionRequest Request Object
type BatchUpdateTestCasesInDiffVersionRequest struct {
	Body *[]TestCaseInfo `json:"body,omitempty"`
}

func (o BatchUpdateTestCasesInDiffVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateTestCasesInDiffVersionRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateTestCasesInDiffVersionRequest", string(data)}, " ")
}
