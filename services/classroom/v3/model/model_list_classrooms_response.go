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
type ListClassroomsResponse struct {
	// 课堂列表
	Classrooms []ClassroomCard `json:"classrooms,omitempty"`
	// 课堂总数
	Total *int32 `json:"total,omitempty"`
}

func (o ListClassroomsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListClassroomsResponse", string(data)}, " ")
}
