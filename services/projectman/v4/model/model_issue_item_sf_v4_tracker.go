package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 工作项类型
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
