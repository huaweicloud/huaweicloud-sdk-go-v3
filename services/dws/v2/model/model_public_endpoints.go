package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type PublicEndpoints struct {
	// 公网连接信息

	PublicConnectInfo *string `json:"public_connect_info,omitempty"`
	// 公网JDBC URL

	JdbcUrl *string `json:"jdbc_url,omitempty"`
}

func (o PublicEndpoints) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicEndpoints struct{}"
	}

	return strings.Join([]string{"PublicEndpoints", string(data)}, " ")
}
