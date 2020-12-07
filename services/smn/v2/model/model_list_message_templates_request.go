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
type ListMessageTemplatesRequest struct {
	Offset              *int32  `json:"offset,omitempty"`
	Limit               *int32  `json:"limit,omitempty"`
	MessageTemplateName *string `json:"message_template_name,omitempty"`
	Protocol            string  `json:"protocol"`
}

func (o ListMessageTemplatesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListMessageTemplatesRequest", string(data)}, " ")
}
