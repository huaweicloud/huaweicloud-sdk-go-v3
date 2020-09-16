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

// Response Object
type ListJobsResponse struct {
	// 作业列表
	Jobs *[]JobCard `json:"jobs,omitempty"`
	// 作业总数
	Total *int32 `json:"total,omitempty"`
}

func (o ListJobsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListJobsResponse", string(data)}, " ")
}
