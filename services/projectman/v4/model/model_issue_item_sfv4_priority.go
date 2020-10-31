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

// 工作项优先级
type IssueItemSfv4Priority struct {
	// 优先级id
	Id *int32 `json:"id,omitempty"`
	// 优先级
	Name *string `json:"name,omitempty"`
}

func (o IssueItemSfv4Priority) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"IssueItemSfv4Priority", string(data)}, " ")
}
