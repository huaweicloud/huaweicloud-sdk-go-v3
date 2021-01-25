package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowAccountStatusRequest struct {
}

func (o ShowAccountStatusRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowAccountStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowAccountStatusRequest", string(data)}, " ")
}
