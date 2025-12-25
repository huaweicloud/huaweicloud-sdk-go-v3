package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteModuleResponse Response Object
type DeleteModuleResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteModuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteModuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteModuleResponse", string(data)}, " ")
}
