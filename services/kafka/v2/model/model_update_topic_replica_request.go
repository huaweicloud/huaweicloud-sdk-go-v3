package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateTopicReplicaRequest struct {
	InstanceId string           `json:"instance_id"`
	Topic      string           `json:"topic"`
	Body       *ResetReplicaReq `json:"body,omitempty"`
}

func (o UpdateTopicReplicaRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateTopicReplicaRequest struct{}"
	}

	return strings.Join([]string{"UpdateTopicReplicaRequest", string(data)}, " ")
}
