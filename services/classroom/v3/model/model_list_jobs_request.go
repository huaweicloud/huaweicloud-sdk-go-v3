/*
    * Classroom
    *
    * devcloud classedge api
    *
*/

package model

// Request Object
type ListJobsRequest struct {
	SourceFrom string `json:"source_from"`
	SourceId string `json:"source_id"`
	Offset int32 `json:"offset,omitempty"`
	Limit int32 `json:"limit,omitempty"`
}
