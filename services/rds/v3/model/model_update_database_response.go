package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateDatabaseResponse struct {
	// 操作结果。

	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDatabaseResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDatabaseResponse struct{}"
	}

	return strings.Join([]string{"UpdateDatabaseResponse", string(data)}, " ")
}
