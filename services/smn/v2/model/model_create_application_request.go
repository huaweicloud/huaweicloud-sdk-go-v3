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
type CreateApplicationRequest struct {
	Body *CreateApplicationRequestBody `json:"body,omitempty"`
}

func (o CreateApplicationRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateApplicationRequest", string(data)}, " ")
}
