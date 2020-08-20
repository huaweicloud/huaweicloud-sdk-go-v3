/*
 * Classroom
 *
 * devcloud classedge api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowJobDetailRequest struct {
	JobId string `json:"job_id"`
}

func (o ShowJobDetailRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowJobDetailRequest", string(data)}, " ")
}
