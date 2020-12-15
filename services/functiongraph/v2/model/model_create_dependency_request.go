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
type CreateDependencyRequest struct {
	Body *CreateDependencyRequestBody `json:"body,omitempty"`
}

func (o CreateDependencyRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateDependencyRequest", string(data)}, " ")
}
