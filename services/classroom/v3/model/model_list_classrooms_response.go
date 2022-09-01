package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListClassroomsResponse struct {

	// 课堂列表
	Classrooms *[]ClassroomCard `json:"classrooms,omitempty" xml:"classrooms"`

	// 课堂总数
	Total          *int32 `json:"total,omitempty" xml:"total"`
	HttpStatusCode int    `json:"-"`
}

func (o ListClassroomsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClassroomsResponse struct{}"
	}

	return strings.Join([]string{"ListClassroomsResponse", string(data)}, " ")
}
