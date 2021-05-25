package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowBigkeyAutoscanConfigRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
}

func (o ShowBigkeyAutoscanConfigRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowBigkeyAutoscanConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowBigkeyAutoscanConfigRequest", string(data)}, " ")
}
