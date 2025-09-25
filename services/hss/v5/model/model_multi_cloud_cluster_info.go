package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MultiCloudClusterInfo 多云集群的集群信息
type MultiCloudClusterInfo struct {

	// 集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 集群服务商
	Provider *string `json:"provider,omitempty"`

	// 集群apiserver地址
	Server *string `json:"server,omitempty"`

	// 镜像仓地址
	ImageRepo *string `json:"image_repo,omitempty"`

	// **参数解释** anp-agent的连接状态 **取值范围**   - not_connect：未连接。   - connect_success：连接成功。   - connect_fail：连接失败。   - connect_disruption：连接中断。
	Status *string `json:"status,omitempty"`

	// anp-agent的版本
	Version *string `json:"version,omitempty"`

	// 当前有效期结束时间
	CurrentExpirationDate *int64 `json:"current_expiration_date,omitempty"`

	// 证书有效期结束时间
	CertificateExpirationDate *int64 `json:"certificate_expiration_date,omitempty"`
}

func (o MultiCloudClusterInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiCloudClusterInfo struct{}"
	}

	return strings.Join([]string{"MultiCloudClusterInfo", string(data)}, " ")
}
