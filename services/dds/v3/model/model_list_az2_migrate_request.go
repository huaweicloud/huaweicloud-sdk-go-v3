package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListAz2MigrateRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o ListAz2MigrateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAz2MigrateRequest struct{}"
	}

	return strings.Join([]string{"ListAz2MigrateRequest", string(data)}, " ")
}
