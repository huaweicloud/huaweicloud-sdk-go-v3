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
type BatchDeleteFunctionTriggersResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteFunctionTriggersResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchDeleteFunctionTriggersResponse", string(data)}, " ")
}
