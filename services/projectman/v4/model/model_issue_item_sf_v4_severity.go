package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueItemSfV4Severity 工作项重要程度
type IssueItemSfV4Severity struct {

	// 重要程度id
	Id *int32 `json:"id,omitempty"`

	// 重要程度
	Name *string `json:"name,omitempty"`
}

func (o IssueItemSfV4Severity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueItemSfV4Severity struct{}"
	}

	return strings.Join([]string{"IssueItemSfV4Severity", string(data)}, " ")
}
