package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIteratorIssueTreeRequest Request Object
type ListIteratorIssueTreeRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 迭代uri
	IteratorId string `json:"iterator_id"`

	Body *IssueTreeInfo `json:"body,omitempty"`
}

func (o ListIteratorIssueTreeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIteratorIssueTreeRequest struct{}"
	}

	return strings.Join([]string{"ListIteratorIssueTreeRequest", string(data)}, " ")
}
