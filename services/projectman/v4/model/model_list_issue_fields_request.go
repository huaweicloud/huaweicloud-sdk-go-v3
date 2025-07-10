package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIssueFieldsRequest Request Object
type ListIssueFieldsRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	// 工作项类型5位id
	CategoryId string `json:"category_id"`

	// 页码
	Page int32 `json:"page"`

	// 每页查询数据数量
	Size int32 `json:"size"`
}

func (o ListIssueFieldsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIssueFieldsRequest struct{}"
	}

	return strings.Join([]string{"ListIssueFieldsRequest", string(data)}, " ")
}
