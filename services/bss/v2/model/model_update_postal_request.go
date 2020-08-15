/*
 * Bss
 *
 * Business Support System API
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Request Object
type UpdatePostalRequest struct {
	XLanguage string           `json:"X-Language,omitempty"`
	Body      *UpdatePostalReq `json:"body,omitempty"`
}

func (o UpdatePostalRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdatePostalRequest", string(data)}, " ")
}
