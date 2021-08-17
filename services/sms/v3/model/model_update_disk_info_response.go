package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateDiskInfoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDiskInfoResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDiskInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdateDiskInfoResponse", string(data)}, " ")
}
