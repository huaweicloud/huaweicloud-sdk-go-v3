package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowResourceDetailCertificateResponseBodyCertificateList struct {

	// 证书名称
	Name *string `json:"name,omitempty"`

	// 证书ID
	Id *string `json:"id,omitempty"`

	// 租户ID
	TenantId *string `json:"tenant_id,omitempty"`

	// 集群ID
	ClusterId *string `json:"cluster_id,omitempty"`

	// 密码服务类型
	ServerType *string `json:"server_type,omitempty"`

	// 证书类型，0：根证书，1：业务证书
	CertificateType *string `json:"certificate_type,omitempty"`

	// 算法类型，0：国密，1：国际
	AlgorithmType *string `json:"algorithm_type,omitempty"`

	// 算法名称
	Algorithm *string `json:"algorithm,omitempty"`

	// 开始时间，UNIX时间戳，单位毫秒
	StartTime *int64 `json:"start_time,omitempty"`

	// 过期时间，UNIX时间戳，单位毫秒
	ExpiredTime *int64 `json:"expired_time,omitempty"`

	// 签发者
	Issuer *string `json:"issuer,omitempty"`

	// 用户
	User *string `json:"user,omitempty"`

	// 证书状态，0：正常，2：过期，3：即将过期
	Status *int32 `json:"status,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`
}

func (o ShowResourceDetailCertificateResponseBodyCertificateList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceDetailCertificateResponseBodyCertificateList struct{}"
	}

	return strings.Join([]string{"ShowResourceDetailCertificateResponseBodyCertificateList", string(data)}, " ")
}
