/*
 * smn
 *
 * SMN Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateMessageTemplateRequest struct {
	Body *CreateMessageTemplateRequestBody `json:"body,omitempty"`
}

func (o CreateMessageTemplateRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateMessageTemplateRequest", string(data)}, " ")
}
