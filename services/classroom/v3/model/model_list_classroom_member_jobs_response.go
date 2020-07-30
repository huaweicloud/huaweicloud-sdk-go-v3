/*
    * Classroom
    *
    * devcloud classedge api
    *
*/

package model

// Response Object
type ListClassroomMemberJobsResponse struct {
	// 课堂下作业列表信息
	Jobs []MemberJobCard `json:"jobs,omitempty"`
	// 学生作业总数
	Total int32 `json:"total,omitempty"`
}
