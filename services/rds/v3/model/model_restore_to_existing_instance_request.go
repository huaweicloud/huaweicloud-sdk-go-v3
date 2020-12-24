/*
 * RDS
 *
 * API v3
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RestoreToExistingInstanceRequest struct {
	XLanguage *string                               `json:"X-Language,omitempty"`
	Body      *RestoreToExistingInstanceRequestBody `json:"body,omitempty"`
}

func (o RestoreToExistingInstanceRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"RestoreToExistingInstanceRequest", string(data)}, " ")
}
