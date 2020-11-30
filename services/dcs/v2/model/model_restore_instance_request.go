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
type RestoreInstanceRequest struct {
	InstanceId string               `json:"instance_id"`
	Body       *RestoreInstanceBody `json:"body,omitempty"`
}

func (o RestoreInstanceRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"RestoreInstanceRequest", string(data)}, " ")
}
