/*
    * Classroom
    *
    * devcloud classedge api
    *
*/

package model

// Request Object
type ListClassroomMembersRequest struct {
	ClassroomId string `json:"classroom_id"`
	Offset int32 `json:"offset,omitempty"`
	Limit int32 `json:"limit,omitempty"`
	Filter string `json:"filter,omitempty"`
}
