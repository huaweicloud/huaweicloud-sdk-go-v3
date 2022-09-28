package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 用户
type UserDto struct {
	Domain *DomainDto `json:"domain,omitempty"`

	// 用户id
	Id *string `json:"id,omitempty"`

	// 用户名
	Name *string `json:"name,omitempty"`
}

func (o UserDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserDto struct{}"
	}

	return strings.Join([]string{"UserDto", string(data)}, " ")
}
