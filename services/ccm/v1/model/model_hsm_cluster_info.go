package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HsmClusterInfo struct {

	// 项目信息
	HsmProject string `json:"hsm_project"`

	// 加密机集群标识符
	HsmClusterId string `json:"hsm_cluster_id"`

	// pem格式证书base64之后的字符串
	HsmCaCert string `json:"hsm_ca_cert"`
}

func (o HsmClusterInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HsmClusterInfo struct{}"
	}

	return strings.Join([]string{"HsmClusterInfo", string(data)}, " ")
}
