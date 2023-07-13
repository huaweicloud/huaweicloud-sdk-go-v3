package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateChildUserNickNameRequestBody struct {

	// 用户昵称
	NickName string `json:"nick_name"`

	// 用户id
	UserId string `json:"user_id"`
}

func (o UpdateChildUserNickNameRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateChildUserNickNameRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateChildUserNickNameRequestBody", string(data)}, " ")
}
