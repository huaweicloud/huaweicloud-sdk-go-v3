package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIssueStatuesRequest Request Object
type ListIssueStatuesRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	// 工作项类型5位id
	CategoryId string `json:"category_id"`
}

func (o ListIssueStatuesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIssueStatuesRequest struct{}"
	}

	return strings.Join([]string{"ListIssueStatuesRequest", string(data)}, " ")
}
