package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type RolePolicy struct {
	// 该权限所依赖的权限。

	Depends *[]PolicyDepends `json:"Depends,omitempty"`
	// 授权语句，描述权限的具体内容。

	Statement []PolicyStatement `json:"Statement"`
	// 权限版本号。 > - 1.0：系统预置的角色。以服务为粒度，提供有限的服务相关角色用于授权。 > - 1.1：策略。IAM最新提供的一种细粒度授权的能力，可以精确到具体服务的操作、资源以及请求条件等。

	Version string `json:"Version"`
}

func (o RolePolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RolePolicy struct{}"
	}

	return strings.Join([]string{"RolePolicy", string(data)}, " ")
}
