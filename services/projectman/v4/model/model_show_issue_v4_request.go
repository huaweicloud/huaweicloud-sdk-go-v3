package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIssueV4Request Request Object
type ShowIssueV4Request struct {

	// devcloud项目的32位id
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
