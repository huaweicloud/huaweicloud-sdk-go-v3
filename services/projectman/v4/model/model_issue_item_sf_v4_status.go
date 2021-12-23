package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 工作项状态
type IssueItemSfV4Status struct {
	// 状态id

	Id *int32 `json:"id,omitempty"`
	// 状态名称

	Name *string `json:"name,omitempty"`
}

func (o IssueItemSfV4Status) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueItemSfV4Status struct{}"
	}

	return strings.Join([]string{"IssueItemSfV4Status", string(data)}, " ")
}
