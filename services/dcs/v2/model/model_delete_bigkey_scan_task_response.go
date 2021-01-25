package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteBigkeyScanTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteBigkeyScanTaskResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteBigkeyScanTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteBigkeyScanTaskResponse", string(data)}, " ")
}
