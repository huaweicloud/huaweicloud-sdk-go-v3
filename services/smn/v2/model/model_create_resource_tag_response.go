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

// Response Object
type CreateResourceTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateResourceTagResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateResourceTagResponse", string(data)}, " ")
}
