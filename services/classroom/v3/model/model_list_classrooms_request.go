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
type ListClassroomsRequest struct {
	Offset    *int32  `json:"offset,omitempty"`
	Limit     *int32  `json:"limit,omitempty"`
	QueryType *string `json:"query_type,omitempty"`
}

func (o ListClassroomsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListClassroomsRequest", string(data)}, " ")
}
