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
type BatchStartBaremetalServersRequest struct {
	Body *OsStartBody `json:"body,omitempty"`
}

func (o BatchStartBaremetalServersRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchStartBaremetalServersRequest", string(data)}, " ")
}
