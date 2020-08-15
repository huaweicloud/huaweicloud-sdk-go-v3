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
type DeletePrivateipRequest struct {
	PrivateipId string `json:"privateip_id"`
}

func (o DeletePrivateipRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeletePrivateipRequest", string(data)}, " ")
}
