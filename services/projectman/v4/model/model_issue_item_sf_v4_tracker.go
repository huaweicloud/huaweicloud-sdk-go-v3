package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueItemSfV4Tracker 工作项类型， 2任务/Task,3缺陷/Bug,5Epic,6Feature,7Story
type IssueItemSfV4Tracker struct {

	// 类型id
	Id *int32 `json:"id,omitempty"`

	// 类型名称
	Name *string `json:"name,omitempty"`
}

func (o IssueItemSfV4Tracker) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueItemSfV4Tracker struct{}"
	}

	return strings.Join([]string{"IssueItemSfV4Tracker", string(data)}, " ")
}
