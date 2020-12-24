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
type ListApplicationAttributesRequest struct {
	ApplicationUrn string `json:"application_urn"`
}

func (o ListApplicationAttributesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListApplicationAttributesRequest", string(data)}, " ")
}
