package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListClassroomMemberJobsRequest struct {
	// 课堂ID

	ClassroomId string `json:"classroom_id"`
	// 用户ID

	MemberId string `json:"member_id"`
	// 信息记录的起始编号

	Offset *int32 `json:"offset,omitempty"`
	// 每页包含的信息记录数

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListClassroomMemberJobsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListClassroomMemberJobsRequest struct{}"
	}

	return strings.Join([]string{"ListClassroomMemberJobsRequest", string(data)}, " ")
}
