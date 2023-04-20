package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchShowUserDetailsResponse struct {

	// 用户信息列表
	Body           *[]UserDto `json:"body,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o BatchShowUserDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowUserDetailsResponse struct{}"
	}

	return strings.Join([]string{"BatchShowUserDetailsResponse", string(data)}, " ")
}
