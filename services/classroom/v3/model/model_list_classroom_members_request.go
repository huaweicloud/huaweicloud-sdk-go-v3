package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListClassroomMembersRequest struct {

	// 课堂ID
	ClassroomId string `json:"classroom_id" xml:"classroom_id"`

	// 信息记录的起始编号
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 每页包含的信息记录数
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 按照条件搜索学生，搜索字段会同时匹配姓名，学号，用户名，班级
	Filter *string `json:"filter,omitempty" xml:"filter"`
}

func (o ListClassroomMembersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClassroomMembersRequest struct{}"
	}

	return strings.Join([]string{"ListClassroomMembersRequest", string(data)}, " ")
}
