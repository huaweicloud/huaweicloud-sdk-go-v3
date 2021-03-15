package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListPluginsRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o ListPluginsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListPluginsRequest struct{}"
	}

	return strings.Join([]string{"ListPluginsRequest", string(data)}, " ")
}
