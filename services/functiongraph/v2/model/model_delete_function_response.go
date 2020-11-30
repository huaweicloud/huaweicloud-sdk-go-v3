/*
 * FunctionGraph
 *
 * API v2
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteFunctionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteFunctionResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteFunctionResponse", string(data)}, " ")
}
