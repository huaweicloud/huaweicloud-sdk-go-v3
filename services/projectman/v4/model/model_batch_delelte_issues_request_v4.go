package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDelelteIssuesRequestV4 struct {
	// 工作项的id

	IssueIds []int32 `json:"issue_ids"`
}

func (o BatchDelelteIssuesRequestV4) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDelelteIssuesRequestV4 struct{}"
	}

	return strings.Join([]string{"BatchDelelteIssuesRequestV4", string(data)}, " ")
}
