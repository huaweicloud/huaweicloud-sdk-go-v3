package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePrivilegesResponse Response Object
type CreatePrivilegesResponse struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePrivilegesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivilegesResponse struct{}"
	}

	return strings.Join([]string{"CreatePrivilegesResponse", string(data)}, " ")
}
