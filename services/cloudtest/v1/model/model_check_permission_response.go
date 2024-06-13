package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckPermissionResponse Response Object
type CheckPermissionResponse struct {

	// 接口调用失败错误码
	Code *string `json:"code,omitempty"`

	// 当前用户权限
	Data map[string]bool `json:"data,omitempty"`

	// 接口调用错误信息
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckPermissionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckPermissionResponse struct{}"
	}

	return strings.Join([]string{"CheckPermissionResponse", string(data)}, " ")
}
