/*
 * SMN
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
type UpdateMessageTemplateRequest struct {
	MessageTemplateId string                            `json:"message_template_id"`
	Body              *UpdateMessageTemplateRequestBody `json:"body,omitempty"`
}

func (o UpdateMessageTemplateRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateMessageTemplateRequest", string(data)}, " ")
}
