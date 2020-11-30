/*
 * ProjectMan
 *
 * devcloud projectman api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListIssueRecordsV4Response struct {
	// 操作记录id
	Id *int32 `json:"id,omitempty"`
	// 创建时间
	CreatedTime *string          `json:"created_time,omitempty"`
	Records     *[]IssueRecordV4 `json:"records,omitempty"`
	// 操作记录总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListIssueRecordsV4Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListIssueRecordsV4Response", string(data)}, " ")
}
