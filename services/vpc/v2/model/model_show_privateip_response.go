/*
 * VPC
 *
 * VPC Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowPrivateipResponse struct {
	Privateip *Privateip `json:"privateip,omitempty"`
}

func (o ShowPrivateipResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowPrivateipResponse", string(data)}, " ")
}
