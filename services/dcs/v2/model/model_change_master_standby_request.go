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
type ChangeMasterStandbyRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o ChangeMasterStandbyRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ChangeMasterStandbyRequest", string(data)}, " ")
}
