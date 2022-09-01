package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type PublicEndpoints struct {

	// 公网连接信息
	PublicConnectInfo *string `json:"public_connect_info,omitempty" xml:"public_connect_info"`

	// 公网JDBC URL，默认格式如下： jdbc:postgresql://<public_connect_info>/<YOUR_DATABASE_name>
	JdbcUrl *string `json:"jdbc_url,omitempty" xml:"jdbc_url"`
}

func (o PublicEndpoints) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicEndpoints struct{}"
	}

	return strings.Join([]string{"PublicEndpoints", string(data)}, " ")
}
