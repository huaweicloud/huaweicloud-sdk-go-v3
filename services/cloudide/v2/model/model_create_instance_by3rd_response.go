package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateInstanceBy3rdResponse struct {
	Result *InstancesResponseInstancesVoResult `json:"result,omitempty"`
	// 状态

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateInstanceBy3rdResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateInstanceBy3rdResponse struct{}"
	}

	return strings.Join([]string{"CreateInstanceBy3rdResponse", string(data)}, " ")
}
