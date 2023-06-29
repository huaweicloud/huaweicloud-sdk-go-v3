package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInitPasswordRequest Request Object
type UpdateInitPasswordRequest struct {

	// 用户id
	UserId string `json:"user_id"`

	Body *ResetPasswordReq `json:"body,omitempty"`
}

func (o UpdateInitPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInitPasswordRequest struct{}"
	}

	return strings.Join([]string{"UpdateInitPasswordRequest", string(data)}, " ")
}
