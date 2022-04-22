package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListIssueRecordsV4Response struct {

	// 操作记录id (已废弃)
	Id *int32 `json:"id,omitempty"`

	// 创建时间 (已废弃)
	CreatedTime *int64 `json:"created_time,omitempty"`

	Records *[]IssueRecordV4 `json:"records,omitempty"`

	// 操作记录总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListIssueRecordsV4Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIssueRecordsV4Response struct{}"
	}

	return strings.Join([]string{"ListIssueRecordsV4Response", string(data)}, " ")
}
