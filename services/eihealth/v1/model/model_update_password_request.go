package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdatePasswordRequest struct {

	// 用户id
	UserId string `json:"user_id"`

	Body *ChangePasswordReq `json:"body,omitempty"`
}

func (o UpdatePasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePasswordRequest struct{}"
	}

	return strings.Join([]string{"UpdatePasswordRequest", string(data)}, " ")
}
