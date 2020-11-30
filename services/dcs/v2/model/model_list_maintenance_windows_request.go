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
type ListMaintenanceWindowsRequest struct {
}

func (o ListMaintenanceWindowsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListMaintenanceWindowsRequest", string(data)}, " ")
}
