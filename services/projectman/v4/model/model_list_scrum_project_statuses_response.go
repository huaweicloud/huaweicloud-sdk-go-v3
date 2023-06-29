package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScrumProjectStatusesResponse Response Object
type ListScrumProjectStatusesResponse struct {

	// 状态总数
	Total *int32 `json:"total,omitempty"`

	// 状态列表
	IssueStatuses  *[]IssueStatus `json:"issue_statuses,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListScrumProjectStatusesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScrumProjectStatusesResponse struct{}"
	}

	return strings.Join([]string{"ListScrumProjectStatusesResponse", string(data)}, " ")
}
