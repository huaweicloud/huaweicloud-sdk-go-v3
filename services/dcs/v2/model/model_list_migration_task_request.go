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

// Request Object
type ListMigrationTaskRequest struct {
	Offset *int32  `json:"offset,omitempty"`
	Limit  *int32  `json:"limit,omitempty"`
	Name   *string `json:"name,omitempty"`
}

func (o ListMigrationTaskRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListMigrationTaskRequest", string(data)}, " ")
}
