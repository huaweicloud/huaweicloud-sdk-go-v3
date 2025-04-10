package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserName 用户名
type UserName struct {
}

func (o UserName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserName struct{}"
	}

	return strings.Join([]string{"UserName", string(data)}, " ")
}
