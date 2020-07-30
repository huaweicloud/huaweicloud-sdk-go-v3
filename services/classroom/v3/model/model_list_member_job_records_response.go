/*
    * Classroom
    *
    * devcloud classedge api
    *
*/

package model

// Response Object
type ListMemberJobRecordsResponse struct {
	// 习题提交列表信息
	Records []JobRecords `json:"records,omitempty"`
	// 习题提交总次数
	Total int32 `json:"total,omitempty"`
}
