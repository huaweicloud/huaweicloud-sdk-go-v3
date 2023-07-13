package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SourceDb 源数据库信息。
type SourceDb struct {

	// 用户名。
	UserName string `json:"user_name"`

	// 连接字符串。
	ConnectionString *string `json:"connection_string,omitempty"`

	// 源数据库类型。
	SourceDbType string `json:"source_db_type"`

	// service名称。
	ServiceName string `json:"service_name"`

	// ip。
	Ip *string `json:"ip,omitempty"`

	// port。
	Port *string `json:"port,omitempty"`
}

func (o SourceDb) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourceDb struct{}"
	}

	return strings.Join([]string{"SourceDb", string(data)}, " ")
}
