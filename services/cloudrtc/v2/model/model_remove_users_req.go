package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RemoveUsersReq struct {

	// 用户ID列表
	UserIds []string `json:"user_ids"`
}

func (o RemoveUsersReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveUsersReq struct{}"
	}

	return strings.Join([]string{"RemoveUsersReq", string(data)}, " ")
}
