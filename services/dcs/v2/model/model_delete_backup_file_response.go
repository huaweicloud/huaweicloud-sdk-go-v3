/*
 * DCS
 *
 * DCS V2版本API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteBackupFileResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteBackupFileResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteBackupFileResponse", string(data)}, " ")
}
