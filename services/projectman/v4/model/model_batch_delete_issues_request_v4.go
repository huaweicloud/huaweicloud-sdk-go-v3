package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteIssuesRequestV4 struct {

	// 工作项的id
	IssueIds []int32 `json:"issue_ids"`
}

func (o BatchDeleteIssuesRequestV4) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteIssuesRequestV4 struct{}"
	}

	return strings.Join([]string{"BatchDeleteIssuesRequestV4", string(data)}, " ")
}
