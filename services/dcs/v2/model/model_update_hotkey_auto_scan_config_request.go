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
type UpdateHotkeyAutoScanConfigRequest struct {
	InstanceId string                 `json:"instance_id"`
	Body       *AutoscanConfigRequest `json:"body,omitempty"`
}

func (o UpdateHotkeyAutoScanConfigRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateHotkeyAutoScanConfigRequest", string(data)}, " ")
}
