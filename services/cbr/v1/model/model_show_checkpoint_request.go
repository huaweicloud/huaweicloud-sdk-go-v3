package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowCheckpointRequest struct {
	// 还原点ID

	CheckpointId string `json:"checkpoint_id"`
}

func (o ShowCheckpointRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowCheckpointRequest struct{}"
	}

	return strings.Join([]string{"ShowCheckpointRequest", string(data)}, " ")
}
