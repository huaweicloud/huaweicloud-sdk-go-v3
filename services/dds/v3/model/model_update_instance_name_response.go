/*
 * DDS
 *
 * API v3
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateInstanceNameResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateInstanceNameResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateInstanceNameResponse", string(data)}, " ")
}
