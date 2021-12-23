package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 工作项优先级
type IssueItemSfV4Priority struct {
	// 优先级id

	Id *int32 `json:"id,omitempty"`
	// 优先级

	Name *string `json:"name,omitempty"`
}

func (o IssueItemSfV4Priority) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueItemSfV4Priority struct{}"
	}

	return strings.Join([]string{"IssueItemSfV4Priority", string(data)}, " ")
}
