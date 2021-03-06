package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteBackupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteBackupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteBackupResponse struct{}"
	}

	return strings.Join([]string{"DeleteBackupResponse", string(data)}, " ")
}
