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
type DeleteReplicationRequest struct {
	InstanceId string `json:"instance_id"`
	GroupId    string `json:"group_id"`
	NodeId     string `json:"node_id"`
}

func (o DeleteReplicationRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteReplicationRequest struct{}"
	}

	return strings.Join([]string{"DeleteReplicationRequest", string(data)}, " ")
}
