package model

import (
	"encoding/json"

	"strings"
)

// This is a auto restart Body Object
type RestartInstanceReq struct {
	Restart *RestarInstanceInfo `json:"restart,omitempty"`
}

func (o RestartInstanceReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RestartInstanceReq struct{}"
	}

	return strings.Join([]string{"RestartInstanceReq", string(data)}, " ")
}
