package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMemberHealthCheckJobRequestBody 参数解释：创建后端服务器检测任务请求body体。
type CreateMemberHealthCheckJobRequestBody struct {
	MemberCheck *CreateMemberHealthCheckJobOption `json:"member_check,omitempty"`
}

func (o CreateMemberHealthCheckJobRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMemberHealthCheckJobRequestBody struct{}"
	}

	return strings.Join([]string{"CreateMemberHealthCheckJobRequestBody", string(data)}, " ")
}
