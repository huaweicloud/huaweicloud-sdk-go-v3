package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveIssuesFromIteratorRequest Request Object
type RemoveIssuesFromIteratorRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 迭代uri
	IteratorId string `json:"iterator_id"`

	Body *RemoveIssuesInfo `json:"body,omitempty"`
}

func (o RemoveIssuesFromIteratorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveIssuesFromIteratorRequest struct{}"
	}

	return strings.Join([]string{"RemoveIssuesFromIteratorRequest", string(data)}, " ")
}
