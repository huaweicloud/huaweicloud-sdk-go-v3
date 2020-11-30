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
type ListGroupReplicationInfoRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o ListGroupReplicationInfoRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListGroupReplicationInfoRequest", string(data)}, " ")
}
