package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePermissionSetReqBody 创建权限集的请求体
type CreatePermissionSetReqBody struct {

	// 权限集描述.
	Description *string `json:"description,omitempty"`

	// 权限集名称.
	Name string `json:"name"`

	// 用于在联合身份验证过程中重定向应用程序中的用户
	RelayState *string `json:"relay_state,omitempty"`

	// 应用程序用户会话在ISO-8601标准中有效的时间长度
	SessionDuration *string `json:"session_duration,omitempty"`

	// 要附加到新权限集的标签
	Tags *[]TagDto `json:"tags,omitempty"`
}

func (o CreatePermissionSetReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePermissionSetReqBody struct{}"
	}

	return strings.Join([]string{"CreatePermissionSetReqBody", string(data)}, " ")
}
