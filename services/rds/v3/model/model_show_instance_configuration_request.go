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
type ShowInstanceConfigurationRequest struct {
	XLanguage  *string `json:"X-Language,omitempty"`
	InstanceId string  `json:"instance_id"`
}

func (o ShowInstanceConfigurationRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowInstanceConfigurationRequest", string(data)}, " ")
}
