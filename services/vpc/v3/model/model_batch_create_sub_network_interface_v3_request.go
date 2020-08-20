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
type BatchCreateSubNetworkInterfaceV3Request struct {
	Body *BatchCreateSubNetworkInterfaceV3RequestBody `json:"body,omitempty"`
}

func (o BatchCreateSubNetworkInterfaceV3Request) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchCreateSubNetworkInterfaceV3Request", string(data)}, " ")
}
