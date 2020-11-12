/*
 * CloudPipeline
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
type ShowInstanceStatusRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`
	TaskId    string  `json:"task_id"`
}

func (o ShowInstanceStatusRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowInstanceStatusRequest", string(data)}, " ")
}
