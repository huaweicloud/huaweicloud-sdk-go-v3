package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssuePassDetailsVo 需求通过情况
type IssuePassDetailsVo struct {

	// 统计测试中的需求
	Testing *int32 `json:"testing,omitempty"`

	// 统计已完成的需求
	Finished *int32 `json:"finished,omitempty"`

	// 统计未完成的需求
	NotTested *int32 `json:"not_tested,omitempty"`
}

func (o IssuePassDetailsVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssuePassDetailsVo struct{}"
	}

	return strings.Join([]string{"IssuePassDetailsVo", string(data)}, " ")
}
