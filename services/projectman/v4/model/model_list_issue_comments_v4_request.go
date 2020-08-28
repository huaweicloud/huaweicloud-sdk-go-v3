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

// Request Object
type ListIssueCommentsV4Request struct {
	ProjectId string `json:"project_id"`
	IssueId   int32  `json:"issue_id"`
	Offset    *int32 `json:"offset,omitempty"`
	Limit     *int32 `json:"limit,omitempty"`
}

func (o ListIssueCommentsV4Request) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListIssueCommentsV4Request", string(data)}, " ")
}
