package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataBase 数据库连接信息。
type DataBase struct {

	// 端口。
	Port *string `json:"port,omitempty"`

	// 连接IP。
	Ip *string `json:"ip,omitempty"`

	// 用户名。
	UserName string `json:"user_name"`

	// 服务名。
	ServiceName string `json:"service_name"`

	// 连接字符串。
	ConnectionString *string `json:"connection_string,omitempty"`
}

func (o DataBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataBase struct{}"
	}

	return strings.Join([]string{"DataBase", string(data)}, " ")
}
