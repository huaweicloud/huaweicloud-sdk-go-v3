package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群的内网连接信息。
type Endpoints struct {

	// 内网连接信息。
	ConnectInfo *string `json:"connect_info,omitempty" xml:"connect_info"`

	// 内网JDBC URL，默认格式如下： jdbc:postgresql://<connect_info>/<YOUR_DATABASE_NAME>
	JdbcUrl *string `json:"jdbc_url,omitempty" xml:"jdbc_url"`
}

func (o Endpoints) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Endpoints struct{}"
	}

	return strings.Join([]string{"Endpoints", string(data)}, " ")
}
