package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdateChildUserNickNamesRequestBody struct {
	// 修改的用户列表

	Users []UpdateChildUserNickNameRequestBody `json:"users"`
}

func (o BatchUpdateChildUserNickNamesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateChildUserNickNamesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpdateChildUserNickNamesRequestBody", string(data)}, " ")
}
