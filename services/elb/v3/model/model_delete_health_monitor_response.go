package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteHealthMonitorResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteHealthMonitorResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteHealthMonitorResponse struct{}"
	}

	return strings.Join([]string{"DeleteHealthMonitorResponse", string(data)}, " ")
}
