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
type DeleteDatabaseRequest struct {
	XLanguage  *string `json:"X-Language,omitempty"`
	InstanceId string  `json:"instance_id"`
	DbName     string  `json:"db_name"`
}

func (o DeleteDatabaseRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteDatabaseRequest", string(data)}, " ")
}
