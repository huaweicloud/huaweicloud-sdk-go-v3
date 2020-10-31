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
type ListChildIssuesV4Response struct {
	// 工作项列表
	Issues *[]IssueResponseV4 `json:"issues,omitempty"`
	// 总数
	Total *int32 `json:"total,omitempty"`
}

func (o ListChildIssuesV4Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListChildIssuesV4Response", string(data)}, " ")
}
