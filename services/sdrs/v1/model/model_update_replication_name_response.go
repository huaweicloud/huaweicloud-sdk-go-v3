package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateReplicationNameResponse struct {
	Replication    *ShowReplicationParams `json:"replication,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o UpdateReplicationNameResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateReplicationNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateReplicationNameResponse", string(data)}, " ")
}
