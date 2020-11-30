/*
 * RabbitMQ
 *
 * RabbitMQ Document API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowMaintainWindowsResponse struct {
	// 支持的维护时间窗列表。
	MaintainWindows *[]ShowMaintainWindowsRespMaintainWindows `json:"maintain_windows,omitempty"`
	HttpStatusCode  int                                       `json:"-"`
}

func (o ShowMaintainWindowsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowMaintainWindowsResponse", string(data)}, " ")
}
