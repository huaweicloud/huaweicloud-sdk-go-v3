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
type ListIssuesV4Request struct {
	ProjectId string              `json:"project_id"`
	Body      *ListIssueRequestV4 `json:"body,omitempty"`
}

func (o ListIssuesV4Request) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListIssuesV4Request", string(data)}, " ")
}
