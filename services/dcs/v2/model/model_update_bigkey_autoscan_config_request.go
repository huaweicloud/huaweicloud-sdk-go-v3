package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateBigkeyAutoscanConfigRequest struct {
	InstanceId string                 `json:"instance_id"`
	Body       *AutoscanConfigRequest `json:"body,omitempty"`
}

func (o UpdateBigkeyAutoscanConfigRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateBigkeyAutoscanConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateBigkeyAutoscanConfigRequest", string(data)}, " ")
}
