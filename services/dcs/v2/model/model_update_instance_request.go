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
type UpdateInstanceRequest struct {
	InstanceId string              `json:"instance_id"`
	Body       *ModifyInstanceBody `json:"body,omitempty"`
}

func (o UpdateInstanceRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateInstanceRequest", string(data)}, " ")
}
