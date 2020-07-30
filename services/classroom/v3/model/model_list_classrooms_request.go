/*
    * Classroom
    *
    * devcloud classedge api
    *
*/

package model

// Request Object
type ListClassroomsRequest struct {
	Offset int32 `json:"offset,omitempty"`
	Limit int32 `json:"limit,omitempty"`
	QueryType string `json:"query_type,omitempty"`
}
