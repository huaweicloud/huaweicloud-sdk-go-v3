package model

import (
	"encoding/json"

	"strings"
)

//
type UpdateProjectStatusRequestBody struct {
	Project *UpdateProjectOption `json:"project"`
}

func (o UpdateProjectStatusRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateProjectStatusRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateProjectStatusRequestBody", string(data)}, " ")
}
