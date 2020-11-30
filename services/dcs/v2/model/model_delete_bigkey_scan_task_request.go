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
type DeleteBigkeyScanTaskRequest struct {
	InstanceId string `json:"instance_id"`
	BigkeyId   string `json:"bigkey_id"`
}

func (o DeleteBigkeyScanTaskRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteBigkeyScanTaskRequest", string(data)}, " ")
}
