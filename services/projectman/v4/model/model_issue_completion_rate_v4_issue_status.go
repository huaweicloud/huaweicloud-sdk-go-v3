package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 工作项不同状态下的数量
type IssueCompletionRateV4IssueStatus struct {

	// 已关闭的工作项
	ClosedNum *int32 `json:"closed_num,omitempty" xml:"closed_num"`

	// 新建的工作项
	NewNum *int32 `json:"new_num,omitempty" xml:"new_num"`

	// 进行中的工作项数目
	ProcessNum *int32 `json:"process_num,omitempty" xml:"process_num"`

	// 已经拒绝的工作项
	RejectedNum *int32 `json:"rejected_num,omitempty" xml:"rejected_num"`

	// 已经解决的工作项
	SolvedNum *int32 `json:"solved_num,omitempty" xml:"solved_num"`

	// 测试中的工作项
	TestNum *int32 `json:"test_num,omitempty" xml:"test_num"`
}

func (o IssueCompletionRateV4IssueStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueCompletionRateV4IssueStatus struct{}"
	}

	return strings.Join([]string{"IssueCompletionRateV4IssueStatus", string(data)}, " ")
}
