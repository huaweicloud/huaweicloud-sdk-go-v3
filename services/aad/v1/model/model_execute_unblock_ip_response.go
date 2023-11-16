package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteUnblockIpResponse Response Object
type ExecuteUnblockIpResponse struct {

	// 业务错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 描述信息
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteUnblockIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteUnblockIpResponse struct{}"
	}

	return strings.Join([]string{"ExecuteUnblockIpResponse", string(data)}, " ")
}
