package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateUserResponse struct {
	// DDM实例帐号名称。

	Name           *string `json:"name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateUserResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateUserResponse struct{}"
	}

	return strings.Join([]string{"UpdateUserResponse", string(data)}, " ")
}
