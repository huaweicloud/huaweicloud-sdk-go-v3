/*
    * Classroom
    *
    * devcloud classedge api
    *
*/

package model

// Response Object
type ListClassroomsResponse struct {
	// 课堂列表
	Classrooms []ClassroomCard `json:"classrooms,omitempty"`
	// 课堂总数
	Total int32 `json:"total,omitempty"`
}
