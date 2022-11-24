package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListProjectIssuesRecordsV4Response struct {

	// 历史记录
	Records *[]IssueAttrHistoryRecord `json:"records,omitempty"`

	// 总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListProjectIssuesRecordsV4Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectIssuesRecordsV4Response struct{}"
	}

	return strings.Join([]string{"ListProjectIssuesRecordsV4Response", string(data)}, " ")
}
