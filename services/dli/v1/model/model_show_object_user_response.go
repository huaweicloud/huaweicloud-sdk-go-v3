package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowObjectUserResponse Response Object
type ShowObjectUserResponse struct {

	// 成功标识
	IsSuccess *bool `json:"is_success,omitempty"`

	// 响应信息
	Message *string `json:"message,omitempty"`

	// 对象名称
	ObjectName *string `json:"object_name,omitempty"`

	// 对象类型
	ObjectType *string `json:"object_type,omitempty"`

	// 对象类型
	Count *int32 `json:"count,omitempty"`

	// 权限信息
	Privileges     *[]ShowDatabaseUsersPrivilege `json:"privileges,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ShowObjectUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowObjectUserResponse struct{}"
	}

	return strings.Join([]string{"ShowObjectUserResponse", string(data)}, " ")
}
