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
type DeleteApplicationRequest struct {
	ApplicationUrn string `json:"application_urn"`
}

func (o DeleteApplicationRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteApplicationRequest", string(data)}, " ")
}
