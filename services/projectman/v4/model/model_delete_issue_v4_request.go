package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIssueV4Request Request Object
type DeleteIssueV4Request struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	// 工作项id
	IssueId int32 `json:"issue_id"`
}

func (o DeleteIssueV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIssueV4Request struct{}"
	}

	return strings.Join([]string{"DeleteIssueV4Request", string(data)}, " ")
}
