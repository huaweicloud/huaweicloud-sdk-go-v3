/*
 * RDS
 *
 * API v3
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteConfigurationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteConfigurationResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteConfigurationResponse", string(data)}, " ")
}
