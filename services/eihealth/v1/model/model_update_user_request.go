package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserRequest Request Object
type UpdateUserRequest struct {

	// 用户id
	UserId string `json:"user_id"`

	Body *UpdateUserReq `json:"body,omitempty"`
}

func (o UpdateUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserRequest struct{}"
	}

	return strings.Join([]string{"UpdateUserRequest", string(data)}, " ")
}
