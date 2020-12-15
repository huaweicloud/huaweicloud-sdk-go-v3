/*
 * BMS
 *
 * BMS Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowJobInfosRequest struct {
	JobId string `json:"job_id"`
}

func (o ShowJobInfosRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowJobInfosRequest", string(data)}, " ")
}
