package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteVaultTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVaultTagResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteVaultTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteVaultTagResponse", string(data)}, " ")
}
