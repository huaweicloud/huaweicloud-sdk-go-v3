package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteApplicationRequest struct {
	// 应用ID。

	ApplicationId string `json:"application_id"`
}

func (o DeleteApplicationRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteApplicationRequest struct{}"
	}

	return strings.Join([]string{"DeleteApplicationRequest", string(data)}, " ")
}
