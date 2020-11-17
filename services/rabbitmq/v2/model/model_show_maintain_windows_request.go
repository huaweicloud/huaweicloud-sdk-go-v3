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

// Request Object
type ShowMaintainWindowsRequest struct {
}

func (o ShowMaintainWindowsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowMaintainWindowsRequest", string(data)}, " ")
}
