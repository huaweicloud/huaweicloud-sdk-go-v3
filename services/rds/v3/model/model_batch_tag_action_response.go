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
type BatchTagActionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchTagActionResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchTagActionResponse", string(data)}, " ")
}
