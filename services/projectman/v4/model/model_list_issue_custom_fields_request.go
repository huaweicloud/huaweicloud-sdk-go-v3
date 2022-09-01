package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListIssueCustomFieldsRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id" xml:"project_id"`

	Body *ListIssueCustomFieldsRequestBody `json:"body,omitempty" xml:"body"`
}

func (o ListIssueCustomFieldsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIssueCustomFieldsRequest struct{}"
	}

	return strings.Join([]string{"ListIssueCustomFieldsRequest", string(data)}, " ")
}
