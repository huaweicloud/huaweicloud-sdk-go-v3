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

// Request Object
type CreateDatabaseRoleRequest struct {
	InstanceId string                         `json:"instance_id"`
	Body       *CreateDatabaseRoleRequestBody `json:"body,omitempty"`
}

func (o CreateDatabaseRoleRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateDatabaseRoleRequest", string(data)}, " ")
}
