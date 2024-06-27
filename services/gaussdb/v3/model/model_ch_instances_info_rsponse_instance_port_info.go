package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChInstancesInfoRsponseInstancePortInfo 端口信息。
type ChInstancesInfoRsponseInstancePortInfo struct {

	// tep端口。取值范围：0~65535。
	TepPort int32 `json:"tep_port"`

	// http端口。取值范围：0~65535。
	HttpPort int32 `json:"http_port"`

	// MySql端口号。取值范围：0~65535。
	MysqlPort int32 `json:"mysql_port"`

	// https端口号。取值范围：0~65535。
	HttpsPort int32 `json:"https_port"`

	// tep安全端口。取值范围：0~65535。
	TepSecurePort int32 `json:"tep_secure_port"`
}

func (o ChInstancesInfoRsponseInstancePortInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChInstancesInfoRsponseInstancePortInfo struct{}"
	}

	return strings.Join([]string{"ChInstancesInfoRsponseInstancePortInfo", string(data)}, " ")
}
