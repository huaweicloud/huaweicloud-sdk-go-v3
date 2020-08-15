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
type ShowJobDetailRequest struct {
	XLanguage string `json:"X-Language,omitempty"`
	JobId     string `json:"job_id"`
}

func (o ShowJobDetailRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowJobDetailRequest", string(data)}, " ")
}
