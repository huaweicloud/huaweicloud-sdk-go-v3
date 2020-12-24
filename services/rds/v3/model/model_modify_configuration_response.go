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
type ModifyConfigurationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ModifyConfigurationResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ModifyConfigurationResponse", string(data)}, " ")
}
