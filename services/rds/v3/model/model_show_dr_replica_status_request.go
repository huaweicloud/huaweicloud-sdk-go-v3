package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDrReplicaStatusRequest struct {
	XLanguage  *string `json:"X-Language,omitempty"`
	InstanceId string  `json:"instance_id"`
}

func (o ShowDrReplicaStatusRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDrReplicaStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowDrReplicaStatusRequest", string(data)}, " ")
}
