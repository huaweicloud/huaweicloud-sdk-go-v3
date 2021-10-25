package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteDisasterRecoveryDrillResponse struct {
	// 成功返回jobId信息

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDisasterRecoveryDrillResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteDisasterRecoveryDrillResponse struct{}"
	}

	return strings.Join([]string{"DeleteDisasterRecoveryDrillResponse", string(data)}, " ")
}
