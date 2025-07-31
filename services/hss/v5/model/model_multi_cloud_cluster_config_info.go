package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MultiCloudClusterConfigInfo 多云集群的配置信息
type MultiCloudClusterConfigInfo struct {

	// apiserver的地址
	Address *string `json:"address,omitempty"`

	// 证书有效期结束时间
	CertificateExpirationDate *int64 `json:"certificate_expiration_date,omitempty"`
}

func (o MultiCloudClusterConfigInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiCloudClusterConfigInfo struct{}"
	}

	return strings.Join([]string{"MultiCloudClusterConfigInfo", string(data)}, " ")
}
