package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoginCbhRequestBody 云堡垒机登录请求体。
type LoginCbhRequestBody struct {

	// 虚拟机id。
	ServerId string `json:"server_id"`
}

func (o LoginCbhRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoginCbhRequestBody struct{}"
	}

	return strings.Join([]string{"LoginCbhRequestBody", string(data)}, " ")
}
