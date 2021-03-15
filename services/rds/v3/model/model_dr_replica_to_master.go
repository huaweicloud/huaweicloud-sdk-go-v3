package model

import (
	"encoding/json"

	"strings"
)

type DrReplicaToMaster struct {
	DrreplicaToMaster *interface{} `json:"drreplica_to_master"`
}

func (o DrReplicaToMaster) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DrReplicaToMaster struct{}"
	}

	return strings.Join([]string{"DrReplicaToMaster", string(data)}, " ")
}
