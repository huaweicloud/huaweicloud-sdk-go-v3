package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDatabaseUserList 数据库用户列表，列表最大长度为50。
type CreateDatabaseUserList struct {

	// 数据库用户名称，数据库用户名称在1到32个字符之间，由字母、数字、下划线组成，不能包含其他特殊字符。
	Name string `json:"name"`

	// 数据库备注,长度不能超过512个字符，不能包含回车和特殊字符!<\"='>&。  该字段只针对新版本的实例生效，必须大于等于指定的内核版本-2.0.13.0，如果不符合内核版本要求，参考升级内核版本升级至最新。
	Comment *string `json:"comment,omitempty"`

	// 数据库用户密码，不能同用户名称相同，非空，至少包含以下字符中的三种：大写字母、小写字母、数字和特殊符号~!@#$%^*-_=+?,()&组成，长度8~32个字符。  建议您输入高强度密码，以提高安全性，防止出现密码被暴力破解等安全风险。
	Password string `json:"password"`

	// 主机IP地址，即允许数据库用户在当前主机连接数据库，默认IP地址为%，表示允许所有地址访问TaurusDB实例。若IP地址为“10.10.10.%”，则表示10.10.10.X的IP地址都可以访问该TaurusDB实例。若您需要添加多个IP地址，请用英文逗号隔开（逗号前后都不能加空格），例如192.168.0.1,172.16.213.9，一次最多创建50个。
	Hosts *[]string `json:"hosts,omitempty"`

	Databases *[]CreateDatabaseList `json:"databases,omitempty"`
}

func (o CreateDatabaseUserList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseUserList struct{}"
	}

	return strings.Join([]string{"CreateDatabaseUserList", string(data)}, " ")
}
