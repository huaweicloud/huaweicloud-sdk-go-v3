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
type ShowClassroomDetailRequest struct {
	ClassroomId string `json:"classroom_id"`
}

func (o ShowClassroomDetailRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowClassroomDetailRequest", string(data)}, " ")
}
