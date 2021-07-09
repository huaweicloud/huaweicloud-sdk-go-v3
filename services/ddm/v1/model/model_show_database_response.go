package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowDatabaseResponse struct {
	Database       *GetDatabaseResponseBean `json:"database,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowDatabaseResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDatabaseResponse struct{}"
	}

	return strings.Join([]string{"ShowDatabaseResponse", string(data)}, " ")
}
