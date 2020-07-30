/*
    * Classroom
    *
    * devcloud classedge api
    *
*/

package model

// Response Object
type ListClassroomMembersResponse struct {
	// 课堂成员列表
	Members []ClassroomMember `json:"members,omitempty"`
	// 课堂成员总数
	Total int32 `json:"total,omitempty"`
}
