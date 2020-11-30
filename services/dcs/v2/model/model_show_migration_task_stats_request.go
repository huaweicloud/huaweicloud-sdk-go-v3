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
type ShowMigrationTaskStatsRequest struct {
	TaskId string `json:"task_id"`
}

func (o ShowMigrationTaskStatsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowMigrationTaskStatsRequest", string(data)}, " ")
}
