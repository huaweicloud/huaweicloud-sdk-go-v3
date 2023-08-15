package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIssueResponseV4Tracker 工作项类型 2任务/Task,3缺陷/Bug,5Epic,6Feature,7Story
type CreateIssueResponseV4Tracker struct {

	// 类型id
	Id *int32 `json:"id,omitempty"`

	// 类型名称
	Name *string `json:"name,omitempty"`
}

func (o CreateIssueResponseV4Tracker) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIssueResponseV4Tracker struct{}"
	}

	return strings.Join([]string{"CreateIssueResponseV4Tracker", string(data)}, " ")
}
