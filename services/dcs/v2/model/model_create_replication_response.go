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
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateReplicationResponse struct{}"
	}

	return strings.Join([]string{"CreateReplicationResponse", string(data)}, " ")
}
