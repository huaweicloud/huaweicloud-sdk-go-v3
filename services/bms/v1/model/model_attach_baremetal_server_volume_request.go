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
type AttachBaremetalServerVolumeRequest struct {
	ServerId string            `json:"server_id"`
	Body     *AttachVolumeBody `json:"body,omitempty"`
}

func (o AttachBaremetalServerVolumeRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"AttachBaremetalServerVolumeRequest", string(data)}, " ")
}
