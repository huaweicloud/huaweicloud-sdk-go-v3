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
type ListConfigurationsResponse struct {
	Configurations *[]ConfigurationSummary `json:"configurations,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListConfigurationsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListConfigurationsResponse", string(data)}, " ")
}
