package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 最后更新时的用户
type LastUpdatedUser struct {

	// 用户ID
	UserId *string `json:"user_id,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`
}

func (o LastUpdatedUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LastUpdatedUser struct{}"
	}

	return strings.Join([]string{"LastUpdatedUser", string(data)}, " ")
}
