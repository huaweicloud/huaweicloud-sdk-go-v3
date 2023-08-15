package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIssueV4Request Request Object
type UpdateIssueV4Request struct {

	// devcloud项目的32位id
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
