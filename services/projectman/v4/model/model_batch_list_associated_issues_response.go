package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchListAssociatedIssuesResponse Response Object
type BatchListAssociatedIssuesResponse struct {

	// 关联的工作项列表
	Issues *[]BatchAssociatedIssue `json:"issues,omitempty"`

	// 总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchListAssociatedIssuesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListAssociatedIssuesResponse struct{}"
	}

	return strings.Join([]string{"BatchListAssociatedIssuesResponse", string(data)}, " ")
}
