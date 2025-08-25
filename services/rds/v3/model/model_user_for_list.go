package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserForList 数据库用户信息。
type UserForList struct {

	// 数据库用户名称。
	Name string `json:"name"`

	// 数据库及其权限。
	Databases []DatabaseWithPrivilegeObject `json:"databases"`

	// 授权用户登录主机IP列表 • 若IP地址为%，则表示允许所有地址访问MySQL实例。 • 若IP地址为“10.10.10.%”，则表示10.10.10.X的IP地址都可以访问该MySQL实例。 • 支持添加多个IP地址。
	Hosts []string `json:"hosts"`

	// 数据库用户备注
	Comment string `json:"comment"`
}

func (o UserForList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserForList struct{}"
	}

	return strings.Join([]string{"UserForList", string(data)}, " ")
}
