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

// Request Object
type ListFunctionsRequest struct {
	Marker   *string `json:"marker,omitempty"`
	Maxitems *string `json:"maxitems,omitempty"`
}

func (o ListFunctionsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListFunctionsRequest", string(data)}, " ")
}
