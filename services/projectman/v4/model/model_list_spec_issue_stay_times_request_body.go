package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListSpecIssueStayTimesRequestBody struct {

	// 项目uuid
	ProjectId *string `json:"project_id,omitempty"`

	// 工作项id字符串列表
	IssueIds *[]string `json:"issue_ids,omitempty"`
}

func (o ListSpecIssueStayTimesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSpecIssueStayTimesRequestBody struct{}"
	}

	return strings.Join([]string{"ListSpecIssueStayTimesRequestBody", string(data)}, " ")
}
