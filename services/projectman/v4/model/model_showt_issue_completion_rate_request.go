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
type ShowtIssueCompletionRateRequest struct {
	ProjectId string `json:"project_id"`
}

func (o ShowtIssueCompletionRateRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowtIssueCompletionRateRequest", string(data)}, " ")
}
