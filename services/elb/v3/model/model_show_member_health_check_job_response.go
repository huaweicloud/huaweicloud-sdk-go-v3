package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMemberHealthCheckJobResponse Response Object
type ShowMemberHealthCheckJobResponse struct {
	MemberCheck *MemberCheckJobInfo `json:"member_check,omitempty"`

	// 参数解释：请求ID。  注：自动生成 。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowMemberHealthCheckJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMemberHealthCheckJobResponse struct{}"
	}

	return strings.Join([]string{"ShowMemberHealthCheckJobResponse", string(data)}, " ")
}
