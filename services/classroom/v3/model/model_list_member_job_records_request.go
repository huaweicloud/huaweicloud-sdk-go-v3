package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMemberJobRecordsRequest Request Object
type ListMemberJobRecordsRequest struct {

	// 作业ID
	JobId string `json:"job_id"`

	// 习题ID
	ExerciseId string `json:"exercise_id"`

	// 用户ID
	MemberId string `json:"member_id"`

	// 信息记录的起始编号
	Offset *int32 `json:"offset,omitempty"`

	// 每页包含的信息记录数
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListMemberJobRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMemberJobRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListMemberJobRecordsRequest", string(data)}, " ")
}
