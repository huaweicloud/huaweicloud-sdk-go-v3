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

// Request Object
type ShowPrivateipRequest struct {
	PrivateipId string `json:"privateip_id"`
}

func (o ShowPrivateipRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowPrivateipRequest", string(data)}, " ")
}
