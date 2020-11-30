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
type BatchStopBaremetalServersRequest struct {
	Body *OsStopBody `json:"body,omitempty"`
}

func (o BatchStopBaremetalServersRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchStopBaremetalServersRequest", string(data)}, " ")
}
