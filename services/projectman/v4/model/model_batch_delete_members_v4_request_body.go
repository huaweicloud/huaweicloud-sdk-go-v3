package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteMembersV4RequestBody struct {

	// 用户id
	UserIds []string `json:"user_ids"`
}

func (o BatchDeleteMembersV4RequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteMembersV4RequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteMembersV4RequestBody", string(data)}, " ")
}
