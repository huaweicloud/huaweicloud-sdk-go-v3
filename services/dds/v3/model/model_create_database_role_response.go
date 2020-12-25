/*
 * DDS
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
type CreateDatabaseRoleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateDatabaseRoleResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateDatabaseRoleResponse", string(data)}, " ")
}
