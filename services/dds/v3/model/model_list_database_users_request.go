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
type ListDatabaseUsersRequest struct {
	InstanceId string  `json:"instance_id"`
	UserName   *string `json:"user_name,omitempty"`
	DbName     *string `json:"db_name,omitempty"`
	Offset     *int32  `json:"offset,omitempty"`
	Limit      *int32  `json:"limit,omitempty"`
}

func (o ListDatabaseUsersRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListDatabaseUsersRequest", string(data)}, " ")
}
