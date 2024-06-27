package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTestCasesByIssueResponse Response Object
type ListTestCasesByIssueResponse struct {

	// 用例状态
	CaseStatusMap map[string]int32 `json:"case_status_map,omitempty"`

	// 新建态
	NewCreate *int32 `json:"new_create,omitempty"`

	// 设计态
	Designing *int32 `json:"designing,omitempty"`

	// 完成态
	Finished *int32 `json:"finished,omitempty"`

	// 测试态
	Testing *int32 `json:"testing,omitempty"`

	// 需求关联的用例数量
	TestCaseNum *int32 `json:"test_case_num,omitempty"`

	// 用例详情
	Testcases *[]TestCaseVo `json:"testcases,omitempty"`

	// 用例总数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTestCasesByIssueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTestCasesByIssueResponse struct{}"
	}

	return strings.Join([]string{"ListTestCasesByIssueResponse", string(data)}, " ")
}
