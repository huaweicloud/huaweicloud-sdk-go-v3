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
type ShowBaremetalServerInterfaceAttachmentsRequest struct {
	ServerId string `json:"server_id"`
}

func (o ShowBaremetalServerInterfaceAttachmentsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowBaremetalServerInterfaceAttachmentsRequest", string(data)}, " ")
}
