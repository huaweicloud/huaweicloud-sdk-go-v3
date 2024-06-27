package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTestcasesByProjectIssuesRelationResponse Response Object
type ListTestcasesByProjectIssuesRelationResponse struct {

	// 用例详情
	Testcases *[]IssuesRelationTestCaseVo `json:"testcases,omitempty"`

	// 用例总数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 页码
	PageNo *int32 `json:"page_no,omitempty"`

	// 每页数量
	PageSize       *int32 `json:"page_size,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTestcasesByProjectIssuesRelationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTestcasesByProjectIssuesRelationResponse struct{}"
	}

	return strings.Join([]string{"ListTestcasesByProjectIssuesRelationResponse", string(data)}, " ")
}
