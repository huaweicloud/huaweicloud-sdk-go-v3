package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMemberHealthCheckJobRequest Request Object
type CreateMemberHealthCheckJobRequest struct {

	// 参数解释：后端服务器ID。
	MemberId string `json:"member_id"`

	Body *CreateMemberHealthCheckJobRequestBody `json:"body,omitempty"`
}

func (o CreateMemberHealthCheckJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMemberHealthCheckJobRequest struct{}"
	}

	return strings.Join([]string{"CreateMemberHealthCheckJobRequest", string(data)}, " ")
}
