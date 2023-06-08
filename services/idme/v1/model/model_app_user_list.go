package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AppUserList struct {

	// 用户ID。
	Id *string `json:"id,omitempty"`

	// 用户名。
	Name *string `json:"name,omitempty"`
}

func (o AppUserList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppUserList struct{}"
	}

	return strings.Join([]string{"AppUserList", string(data)}, " ")
}
