package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAssociatedTestCasesResponse struct {

	// 关联的测试用例列表
	TestCases *[]AssociatedTestCase `json:"test_cases,omitempty"`

	// 总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAssociatedTestCasesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssociatedTestCasesResponse struct{}"
	}

	return strings.Join([]string{"ListAssociatedTestCasesResponse", string(data)}, " ")
}
