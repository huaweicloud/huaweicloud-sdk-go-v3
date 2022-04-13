package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListClassroomMemberJobsResponse struct {
	// 课堂下作业列表信息

	Jobs *[]MemberJobCard `json:"jobs,omitempty"`
	// 学生作业总数

	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListClassroomMemberJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClassroomMemberJobsResponse struct{}"
	}

	return strings.Join([]string{"ListClassroomMemberJobsResponse", string(data)}, " ")
}
