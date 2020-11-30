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
type CreateBareMetalServersRequest struct {
	Body *CreateBaremetalServersBody `json:"body,omitempty"`
}

func (o CreateBareMetalServersRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateBareMetalServersRequest", string(data)}, " ")
}
