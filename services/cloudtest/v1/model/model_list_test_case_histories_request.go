package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTestCaseHistoriesRequest Request Object
type ListTestCaseHistoriesRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 用例ID，长度11-34位字符（字母和数字）。
	TestcaseId string `json:"testcase_id"`

	Body *ListTestCaseHistoriesRequestBody `json:"body,omitempty"`
}

func (o ListTestCaseHistoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTestCaseHistoriesRequest struct{}"
	}

	return strings.Join([]string{"ListTestCaseHistoriesRequest", string(data)}, " ")
}
