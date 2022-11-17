package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteUserRequest struct {

	// 用户ID。
	UserId string `json:"user_id"`
}

func (o DeleteUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteUserRequest struct{}"
	}

	return strings.Join([]string{"DeleteUserRequest", string(data)}, " ")
}
