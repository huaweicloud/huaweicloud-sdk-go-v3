/*
 * devcloudpipeline
 *
 * devcloud pipeline api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreatePipelineByTemplateRequest struct {
	XLanguage *string       `json:"X-Language,omitempty"`
	Body      *TemplateCddl `json:"body,omitempty"`
}

func (o CreatePipelineByTemplateRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreatePipelineByTemplateRequest", string(data)}, " ")
}
