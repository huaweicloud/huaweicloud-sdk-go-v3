/*
    * Classroom
    *
    * devcloud classedge api
    *
*/

package model

// Response Object
type ListJobsResponse struct {
	// 作业列表
	Jobs []JobCard `json:"jobs,omitempty"`
	// 作业总数
	Total int32 `json:"total,omitempty"`
}
