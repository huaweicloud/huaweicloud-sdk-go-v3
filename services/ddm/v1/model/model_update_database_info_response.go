package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateDatabaseInfoResponse struct {
	// DDM实例id。

	InstanceId *string `json:"instanceId,omitempty"`
	// 任务ID。

	JobId          *string `json:"jobId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDatabaseInfoResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDatabaseInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdateDatabaseInfoResponse", string(data)}, " ")
}
