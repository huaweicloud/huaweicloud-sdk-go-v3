/*
 * Kafka
 *
 * Kafka Document API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListAvailableZonesRequest struct {
}

func (o ListAvailableZonesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListAvailableZonesRequest", string(data)}, " ")
}
