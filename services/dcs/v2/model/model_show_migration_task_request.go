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
type ShowMigrationTaskRequest struct {
	TaskId string `json:"task_id"`
}

func (o ShowMigrationTaskRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowMigrationTaskRequest", string(data)}, " ")
}
