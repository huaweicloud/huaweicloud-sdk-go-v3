package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteUserResponse struct {
	// DDM实例帐号名称。

	Name           *string `json:"name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteUserResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteUserResponse struct{}"
	}

	return strings.Join([]string{"DeleteUserResponse", string(data)}, " ")
}
