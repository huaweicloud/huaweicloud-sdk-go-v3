package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowIssueV4Request struct {
	// devcloud的项目id

	ProjectId string `json:"project_id"`
	// 工作项id

	IssueId int32 `json:"issue_id"`
}

func (o ShowIssueV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIssueV4Request struct{}"
	}

	return strings.Join([]string{"ShowIssueV4Request", string(data)}, " ")
}
