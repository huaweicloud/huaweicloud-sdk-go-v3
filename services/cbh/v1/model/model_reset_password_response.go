package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ResetPasswordResponse struct {

	// 请求得到的信息
	RequestInfo    *string `json:"request_info,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResetPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetPasswordResponse struct{}"
	}

	return strings.Join([]string{"ResetPasswordResponse", string(data)}, " ")
}
