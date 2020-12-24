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

// Request Object
type ShowConfigurationRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`
	ConfigId  string  `json:"config_id"`
}

func (o ShowConfigurationRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowConfigurationRequest", string(data)}, " ")
}
