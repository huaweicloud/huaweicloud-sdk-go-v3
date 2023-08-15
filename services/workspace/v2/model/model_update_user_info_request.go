package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserInfoRequest Request Object
type UpdateUserInfoRequest struct {

	// 用户ID。
	UserId string `json:"user_id"`

	Body *EditUserReq `json:"body,omitempty"`
}

func (o UpdateUserInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserInfoRequest struct{}"
	}

	return strings.Join([]string{"UpdateUserInfoRequest", string(data)}, " ")
}
