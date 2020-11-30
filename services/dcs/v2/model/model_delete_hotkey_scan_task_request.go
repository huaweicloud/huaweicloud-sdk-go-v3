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
type DeleteHotkeyScanTaskRequest struct {
	InstanceId string `json:"instance_id"`
	HotkeyId   string `json:"hotkey_id"`
}

func (o DeleteHotkeyScanTaskRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteHotkeyScanTaskRequest", string(data)}, " ")
}
