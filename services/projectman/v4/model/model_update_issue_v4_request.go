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
type UpdateIssueV4Request struct {
	ProjectId string          `json:"project_id"`
	IssueId   int32           `json:"issue_id"`
	Body      *IssueRequestV4 `json:"body,omitempty"`
}

func (o UpdateIssueV4Request) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateIssueV4Request", string(data)}, " ")
}
