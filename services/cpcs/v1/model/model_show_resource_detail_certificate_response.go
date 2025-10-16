package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResourceDetailCertificateResponse Response Object
type ShowResourceDetailCertificateResponse struct {

	// 资源名称
	MetricName *string `json:"metric_name,omitempty"`

	// 证书详情列表
	CertificateList *[]ShowResourceDetailCertificateResponseBodyCertificateList `json:"certificate_list,omitempty"`

	// 证书总数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 证书过期数量
	ExpiredCount *int32 `json:"expired_count,omitempty"`

	// 证书即将过期数量
	ExpiringCount  *int32 `json:"expiring_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowResourceDetailCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceDetailCertificateResponse struct{}"
	}

	return strings.Join([]string{"ShowResourceDetailCertificateResponse", string(data)}, " ")
}
