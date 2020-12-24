/*
 * RDS
 *
 * API v3
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListJobInfoResponse struct {
	Job            *GetJobInfoResponseBodyJob `json:"job,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListJobInfoResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListJobInfoResponse", string(data)}, " ")
}
