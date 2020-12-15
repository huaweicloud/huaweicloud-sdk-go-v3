/*
 * BMS
 *
 * BMS Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ChangeBaremetalServerNameRequest struct {
	ServerId string                   `json:"server_id"`
	Body     *ChangeBaremetalNameBody `json:"body,omitempty"`
}

func (o ChangeBaremetalServerNameRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ChangeBaremetalServerNameRequest", string(data)}, " ")
}
