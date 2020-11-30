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
type UpdateSlavePriorityRequest struct {
	InstanceId string        `json:"instance_id"`
	GroupId    string        `json:"group_id"`
	NodeId     string        `json:"node_id"`
	Body       *PriorityBody `json:"body,omitempty"`
}

func (o UpdateSlavePriorityRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateSlavePriorityRequest", string(data)}, " ")
}
