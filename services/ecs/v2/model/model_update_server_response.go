/*
 * ECS
 *
 * ECS Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateServerResponse struct {
	Server         *UpdateServerResult `json:"server,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o UpdateServerResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateServerResponse", string(data)}, " ")
}
