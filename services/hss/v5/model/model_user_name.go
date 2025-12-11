package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserName **参数解释**: 用户名 **取值范围**: 字符长度1-64位
type UserName struct {
}

func (o UserName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserName struct{}"
	}

	return strings.Join([]string{"UserName", string(data)}, " ")
}
