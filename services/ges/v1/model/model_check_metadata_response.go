package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CheckMetadataResponse struct {
	// 系统提示信息，执行成功时，字段可能为空。执行失败时，用于显示错误信息。

	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 系统提示信息，执行成功时，字段可能为空。执行失败时，用于显示错误码。

	ErrorCode      *string `json:"errorCode,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckMetadataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckMetadataResponse struct{}"
	}

	return strings.Join([]string{"CheckMetadataResponse", string(data)}, " ")
}
