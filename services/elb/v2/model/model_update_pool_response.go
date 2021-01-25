package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdatePoolResponse struct {
	Pool           *PoolResp `json:"pool,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o UpdatePoolResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePoolResponse struct{}"
	}

	return strings.Join([]string{"UpdatePoolResponse", string(data)}, " ")
}
