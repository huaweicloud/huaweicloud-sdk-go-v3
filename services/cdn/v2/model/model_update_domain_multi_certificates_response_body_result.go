package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDomainMultiCertificatesResponseBodyResult 匹配的域名的详情
type UpdateDomainMultiCertificatesResponseBodyResult struct {

	// 域名名称
	DomainName *string `json:"domain_name,omitempty"`

	// 执行结果，success，fail
	Status *string `json:"status,omitempty"`

	// 失败原因
	Reason *string `json:"reason,omitempty"`
}

func (o UpdateDomainMultiCertificatesResponseBodyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainMultiCertificatesResponseBodyResult struct{}"
	}

	return strings.Join([]string{"UpdateDomainMultiCertificatesResponseBodyResult", string(data)}, " ")
}
