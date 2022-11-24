package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 迭代
type IssueItemSfV4Iteration struct {

	// 迭代id
	Id *int32 `json:"id,omitempty"`

	// 迭代名
	Name *string `json:"name,omitempty"`
}

func (o IssueItemSfV4Iteration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueItemSfV4Iteration struct{}"
	}

	return strings.Join([]string{"IssueItemSfV4Iteration", string(data)}, " ")
}
