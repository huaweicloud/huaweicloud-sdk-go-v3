/*
    * Classroom
    *
    * devcloud classedge api
    *
*/

package model

// Request Object
type ListClassroomMemberJobsRequest struct {
	ClassroomId string `json:"classroom_id"`
	MemberId string `json:"member_id"`
	Offset int32 `json:"offset,omitempty"`
	Limit int32 `json:"limit,omitempty"`
}
