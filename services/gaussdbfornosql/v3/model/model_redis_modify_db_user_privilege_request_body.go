package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RedisModifyDbUserPrivilegeRequestBody 修改数据库账号权限请求体
type RedisModifyDbUserPrivilegeRequestBody struct {

	// 账号名称。  小于36个字符，以字母开头，仅包含数字、字母、中划线、下划线。
	Name string `json:"name"`

	// 账号权限。  - 取值\"ReadOnly\"：账号为只读权限； - 取值\"ReadWrite\"：账号为读写权限。
	Privilege string `json:"privilege"`

	// 账号授权database列表。 不传值则对账号授权的db不做修改。
	Databases *[]string `json:"databases,omitempty"`
}

func (o RedisModifyDbUserPrivilegeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedisModifyDbUserPrivilegeRequestBody struct{}"
	}

	return strings.Join([]string{"RedisModifyDbUserPrivilegeRequestBody", string(data)}, " ")
}
