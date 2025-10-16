package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VendorCertificateStatistic struct {

	// 总数量
	Count *int32 `json:"count,omitempty"`

	// 租户ID
	TenantId *string `json:"tenant_id,omitempty"`

	// 集群ID
	ClusterId *string `json:"cluster_id,omitempty"`

	// 服务类型
	ServerType *string `json:"server_type,omitempty"`

	// 证书类型
	CertificateType *int32 `json:"certificate_type,omitempty"`

	// 算法类型
	AlgorithmType *int32 `json:"algorithm_type,omitempty"`
}

func (o VendorCertificateStatistic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VendorCertificateStatistic struct{}"
	}

	return strings.Join([]string{"VendorCertificateStatistic", string(data)}, " ")
}
