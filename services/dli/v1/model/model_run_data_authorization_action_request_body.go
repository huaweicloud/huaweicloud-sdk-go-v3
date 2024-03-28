package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunDataAuthorizationActionRequestBody 数据授权的请求body体。
type RunDataAuthorizationActionRequestBody struct {

	// 被赋权的用户名称，该用户将有权访问指定的数据库或数据表，被收回或者更新访问权限。
	UserName string `json:"user_name"`

	// 指定赋权或回收。值为：grant，revoke或update。 说明： 当用户同时拥有grant和revoke权限的时候才有权限使用update操作。
	Action string `json:"action"`

	// 赋权信息。
	Privileges []DataAuthorizationPrivilege `json:"privileges"`
}

func (o RunDataAuthorizationActionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunDataAuthorizationActionRequestBody struct{}"
	}

	return strings.Join([]string{"RunDataAuthorizationActionRequestBody", string(data)}, " ")
}
