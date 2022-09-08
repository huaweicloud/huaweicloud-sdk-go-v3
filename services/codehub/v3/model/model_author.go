package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Author struct {

	// 用户id
	Id *float64 `json:"id,omitempty"`

	// 用户名
	Name *string `json:"name,omitempty"`

	// 用户状态
	State *string `json:"state,omitempty"`

	// 用户iamId
	Username *string `json:"username,omitempty"`
}

func (o Author) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Author struct{}"
	}

	return strings.Join([]string{"Author", string(data)}, " ")
}
