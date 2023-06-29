package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddIssueWorkHoursRequest Request Object
type AddIssueWorkHoursRequest struct {

	// 项目id
	ProjectId string `json:"project_id"`

	// 工作项id
	IssueId int32 `json:"issue_id"`

	Body *AddIssueWorkHoursRequestBody `json:"body,omitempty"`
}

func (o AddIssueWorkHoursRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddIssueWorkHoursRequest struct{}"
	}

	return strings.Join([]string{"AddIssueWorkHoursRequest", string(data)}, " ")
}
