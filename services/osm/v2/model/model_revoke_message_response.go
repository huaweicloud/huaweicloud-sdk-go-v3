package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RevokeMessageResponse Response Object
type RevokeMessageResponse struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RevokeMessageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RevokeMessageResponse struct{}"
	}

	return strings.Join([]string{"RevokeMessageResponse", string(data)}, " ")
}
