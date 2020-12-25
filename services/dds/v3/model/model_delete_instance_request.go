/*
 * DDS
 *
 * API v3
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteInstanceRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o DeleteInstanceRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteInstanceRequest", string(data)}, " ")
}
