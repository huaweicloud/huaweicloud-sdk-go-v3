package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeUserStatusRequest Request Object
type ChangeUserStatusRequest struct {

	// 用户ID。
	UserId string `json:"user_id"`

	Body *OperateUserReq `json:"body,omitempty"`
}

func (o ChangeUserStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeUserStatusRequest struct{}"
	}

	return strings.Join([]string{"ChangeUserStatusRequest", string(data)}, " ")
}
