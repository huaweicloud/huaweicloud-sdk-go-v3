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
type CreateReplicationRequest struct {
	InstanceId string              `json:"instance_id"`
	GroupId    string              `json:"group_id"`
	Body       *AddReplicationBody `json:"body,omitempty"`
}

func (o CreateReplicationRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateReplicationRequest struct{}"
	}

	return strings.Join([]string{"CreateReplicationRequest", string(data)}, " ")
}
