package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowConsoleConfigRequest struct {
}

func (o ShowConsoleConfigRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowConsoleConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowConsoleConfigRequest", string(data)}, " ")
}
