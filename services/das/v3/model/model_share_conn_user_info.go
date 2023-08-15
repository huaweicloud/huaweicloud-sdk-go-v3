package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShareConnUserInfo struct {

	// 用户ID
	UserId string `json:"user_id"`

	// 用户名
	UserName string `json:"user_name"`
}

func (o ShareConnUserInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShareConnUserInfo struct{}"
	}

	return strings.Join([]string{"ShareConnUserInfo", string(data)}, " ")
}
