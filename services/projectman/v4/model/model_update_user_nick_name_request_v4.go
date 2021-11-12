package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 用户昵称
type UpdateUserNickNameRequestV4 struct {
	// 用户昵称

	NickName string `json:"nick_name"`
}

func (o UpdateUserNickNameRequestV4) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserNickNameRequestV4 struct{}"
	}

	return strings.Join([]string{"UpdateUserNickNameRequestV4", string(data)}, " ")
}
