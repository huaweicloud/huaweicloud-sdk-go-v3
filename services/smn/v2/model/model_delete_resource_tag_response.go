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
type DeleteResourceTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteResourceTagResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteResourceTagResponse", string(data)}, " ")
}
