package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UserRequest struct {

	// 用户32位字符串id
	UserId string `json:"user_id"`
}

func (o UserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserRequest struct{}"
	}

	return strings.Join([]string{"UserRequest", string(data)}, " ")
}
