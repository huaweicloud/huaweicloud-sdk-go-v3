package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UserSafeDto struct {

	// 用户id
	Id *int32 `json:"id,omitempty"`

	// 姓名
	Name *string `json:"name,omitempty"`

	// 用户名
	Username *string `json:"username,omitempty"`
}

func (o UserSafeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserSafeDto struct{}"
	}

	return strings.Join([]string{"UserSafeDto", string(data)}, " ")
}
