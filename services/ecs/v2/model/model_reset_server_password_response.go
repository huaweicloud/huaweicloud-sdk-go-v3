/*
 * ecs
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
type ResetServerPasswordResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResetServerPasswordResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ResetServerPasswordResponse", string(data)}, " ")
}
