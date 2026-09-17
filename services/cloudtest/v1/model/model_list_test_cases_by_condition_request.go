package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTestCasesByConditionRequest Request Object
type ListTestCasesByConditionRequest struct {

	// 项目ID
	ProjectUuid string `json:"project_uuid"`

	Body *TestCasesListQueryInfo `json:"body,omitempty"`
}

func (o ListTestCasesByConditionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTestCasesByConditionRequest struct{}"
	}

	return strings.Join([]string{"ListTestCasesByConditionRequest", string(data)}, " ")
}
