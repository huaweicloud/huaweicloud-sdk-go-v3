package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListClassroomMemberJobsRequest struct {

	// 课堂ID
	ClassroomId string `json:"classroom_id" xml:"classroom_id"`

	// 用户ID
	MemberId string `json:"member_id" xml:"member_id"`

	// 信息记录的起始编号
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 每页包含的信息记录数
	Limit *int32 `json:"limit,omitempty" xml:"limit"`
}

func (o ListClassroomMemberJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClassroomMemberJobsRequest struct{}"
	}

	return strings.Join([]string{"ListClassroomMemberJobsRequest", string(data)}, " ")
}
