/*
 * VPC
 *
 * VPC Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdatePortResponse struct {
	Port *Port `json:"port,omitempty"`
}

func (o UpdatePortResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdatePortResponse", string(data)}, " ")
}
