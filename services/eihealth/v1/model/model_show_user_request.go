package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUserRequest Request Object
type ShowUserRequest struct {

	// 用户id
	UserId string `json:"user_id"`
}

func (o ShowUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserRequest struct{}"
	}

	return strings.Join([]string{"ShowUserRequest", string(data)}, " ")
}
