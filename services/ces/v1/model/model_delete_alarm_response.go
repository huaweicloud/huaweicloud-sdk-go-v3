package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteAlarmResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAlarmResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteAlarmResponse struct{}"
	}

	return strings.Join([]string{"DeleteAlarmResponse", string(data)}, " ")
}
