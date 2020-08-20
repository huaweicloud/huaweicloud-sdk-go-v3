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
type DeletePortResponse struct {
}

func (o DeletePortResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeletePortResponse", string(data)}, " ")
}
