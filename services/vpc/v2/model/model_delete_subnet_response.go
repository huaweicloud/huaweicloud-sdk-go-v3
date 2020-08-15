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
type DeleteSubnetResponse struct {
}

func (o DeleteSubnetResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteSubnetResponse", string(data)}, " ")
}
