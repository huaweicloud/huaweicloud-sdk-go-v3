/*
 * DCS
 *
 * DCS V2版本API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateInstanceRequest struct {
	Body *CreateInstanceBody `json:"body,omitempty"`
}

func (o CreateInstanceRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateInstanceRequest", string(data)}, " ")
}
