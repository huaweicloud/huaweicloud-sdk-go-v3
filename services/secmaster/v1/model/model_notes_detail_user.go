package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NotesDetailUser 用户信息
type NotesDetailUser struct {

	// 用户id
	Id *string `json:"id,omitempty"`

	// 用户名称
	Name *string `json:"name,omitempty"`
}

func (o NotesDetailUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotesDetailUser struct{}"
	}

	return strings.Join([]string{"NotesDetailUser", string(data)}, " ")
}
