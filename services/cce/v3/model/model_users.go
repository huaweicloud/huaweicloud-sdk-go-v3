package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Users struct {

	// 当前为固定值“user“。
	Name *string `json:"name,omitempty" xml:"name"`

	User *User `json:"user,omitempty" xml:"user"`
}

func (o Users) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Users struct{}"
	}

	return strings.Join([]string{"Users", string(data)}, " ")
}
