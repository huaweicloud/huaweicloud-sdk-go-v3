package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangePasswordRequest Request Object
type ChangePasswordRequest struct {

	// 用户id
	UserId string `json:"user_id"`

	Body *ChangePasswordReq `json:"body,omitempty"`
}

func (o ChangePasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangePasswordRequest struct{}"
	}

	return strings.Join([]string{"ChangePasswordRequest", string(data)}, " ")
}
