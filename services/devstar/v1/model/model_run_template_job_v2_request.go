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
type RunTemplateJobV2Request struct {
	XLanguage *string          `json:"X-Language,omitempty"`
	Body      *TemplateJobInfo `json:"body,omitempty"`
}

func (o RunTemplateJobV2Request) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"RunTemplateJobV2Request", string(data)}, " ")
}
