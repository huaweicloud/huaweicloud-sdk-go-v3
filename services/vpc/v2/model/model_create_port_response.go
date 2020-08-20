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
type CreatePortResponse struct {
	Port *Port `json:"port,omitempty"`
}

func (o CreatePortResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreatePortResponse", string(data)}, " ")
}
