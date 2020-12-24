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
type EnableConfigurationRequest struct {
	XLanguage *string                    `json:"X-Language,omitempty"`
	ConfigId  string                     `json:"config_id"`
	Body      *ApplyConfigurationRequest `json:"body,omitempty"`
}

func (o EnableConfigurationRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"EnableConfigurationRequest", string(data)}, " ")
}
