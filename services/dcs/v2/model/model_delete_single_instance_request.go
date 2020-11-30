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
type DeleteSingleInstanceRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o DeleteSingleInstanceRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteSingleInstanceRequest", string(data)}, " ")
}
