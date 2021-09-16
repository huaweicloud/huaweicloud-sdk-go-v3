package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteDeploymentResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDeploymentResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteDeploymentResponse struct{}"
	}

	return strings.Join([]string{"DeleteDeploymentResponse", string(data)}, " ")
}
