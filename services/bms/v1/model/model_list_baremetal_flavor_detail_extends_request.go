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
type ListBaremetalFlavorDetailExtendsRequest struct {
	AvailabilityZone *string `json:"availability_zone,omitempty"`
}

func (o ListBaremetalFlavorDetailExtendsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListBaremetalFlavorDetailExtendsRequest", string(data)}, " ")
}
