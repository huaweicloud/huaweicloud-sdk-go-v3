package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStatisticCertificateResponse Response Object
type ShowStatisticCertificateResponse struct {

	// 资源名称
	MetricName *string `json:"metric_name,omitempty"`

	// 证书分布按算法和证书类型统计
	CertificateClassifiedCounts *[]VendorCertificateStatistic `json:"certificate_classified_counts,omitempty"`

	// 证书分布按服务类型统计
	CertificateCountsByServerType *[]VendorCertificateStatistic `json:"certificate_counts_by_server_type,omitempty"`
	HttpStatusCode                int                           `json:"-"`
}

func (o ShowStatisticCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStatisticCertificateResponse struct{}"
	}

	return strings.Join([]string{"ShowStatisticCertificateResponse", string(data)}, " ")
}
