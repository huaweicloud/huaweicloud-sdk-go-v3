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
type DeleteManualBackupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteManualBackupResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteManualBackupResponse", string(data)}, " ")
}
