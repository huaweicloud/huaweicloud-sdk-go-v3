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
type CreateReplicationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateReplicationResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateReplicationResponse", string(data)}, " ")
}
