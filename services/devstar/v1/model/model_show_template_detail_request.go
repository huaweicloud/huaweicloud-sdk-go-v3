/*
 * Devstar
 *
 * Devstar API
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Request Object
type ShowTemplateDetailRequest struct {
	XLanguage  string `json:"X-Language,omitempty"`
	TemplateId string `json:"template_id"`
}

func (o ShowTemplateDetailRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowTemplateDetailRequest", string(data)}, " ")
}
