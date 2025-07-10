package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIssueConfigFieldsRequest Request Object
type ShowIssueConfigFieldsRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	// 工作项类型
	IssueCategory string `json:"issue_category"`
}

func (o ShowIssueConfigFieldsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIssueConfigFieldsRequest struct{}"
	}

	return strings.Join([]string{"ShowIssueConfigFieldsRequest", string(data)}, " ")
}
