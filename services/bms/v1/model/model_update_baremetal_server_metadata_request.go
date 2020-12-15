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
type UpdateBaremetalServerMetadataRequest struct {
	ServerId string    `json:"server_id"`
	Body     *MetaData `json:"body,omitempty"`
}

func (o UpdateBaremetalServerMetadataRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateBaremetalServerMetadataRequest", string(data)}, " ")
}
