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
type UpdateDependencyRequest struct {
	DependId string                       `json:"depend_id"`
	Body     *UpdateDependencyRequestBody `json:"body,omitempty"`
}

func (o UpdateDependencyRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateDependencyRequest", string(data)}, " ")
}
