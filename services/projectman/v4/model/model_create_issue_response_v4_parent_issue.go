package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 父工作项
type CreateIssueResponseV4ParentIssue struct {

	// 父工作项id
	Id *int32 `json:"id,omitempty" xml:"id"`

	// 父工作项
	Name *string `json:"name,omitempty" xml:"name"`
}

func (o CreateIssueResponseV4ParentIssue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIssueResponseV4ParentIssue struct{}"
	}

	return strings.Join([]string{"CreateIssueResponseV4ParentIssue", string(data)}, " ")
}
