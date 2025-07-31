package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserGroupName **参数解释**： 用户组 **取值范围**： 字符长度1-128位
type UserGroupName struct {
}

func (o UserGroupName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserGroupName struct{}"
	}

	return strings.Join([]string{"UserGroupName", string(data)}, " ")
}
