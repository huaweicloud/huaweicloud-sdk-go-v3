package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateIssueV4Request struct {
	// devcloud的项目id

	ProjectId string `json:"project_id"`
	// 工作项id

	IssueId int32 `json:"issue_id"`

	Body *IssueRequestV4 `json:"body,omitempty"`
}

func (o UpdateIssueV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIssueV4Request struct{}"
	}

	return strings.Join([]string{"UpdateIssueV4Request", string(data)}, " ")
}
