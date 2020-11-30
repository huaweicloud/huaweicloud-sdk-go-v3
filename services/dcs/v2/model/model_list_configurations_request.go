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
type ListConfigurationsRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o ListConfigurationsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListConfigurationsRequest", string(data)}, " ")
}
