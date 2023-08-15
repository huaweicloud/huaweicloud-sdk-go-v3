package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEnhancedPrivilegeResponse Response Object
type ShowEnhancedPrivilegeResponse struct {

	// 行请求是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message *string `json:"message,omitempty"`

	// 增强型跨源连接ID，用于标识跨源连接的UUID。
	ConnectionId *string `json:"connection_id,omitempty"`

	// 跨源连接各个授权项目的信息。
	Privileges     *[]ProjectPrivilege `json:"privileges,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowEnhancedPrivilegeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEnhancedPrivilegeResponse struct{}"
	}

	return strings.Join([]string{"ShowEnhancedPrivilegeResponse", string(data)}, " ")
}
