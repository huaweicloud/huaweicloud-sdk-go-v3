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

// Response Object
type ListMaintenanceWindowsResponse struct {
	// 支持的维护时间窗列表。
	MaintainWindows *[]MaintainWindows `json:"maintain_windows,omitempty"`
	HttpStatusCode  int                `json:"-"`
}

func (o ListMaintenanceWindowsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListMaintenanceWindowsResponse", string(data)}, " ")
}
