package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CertInfoResp struct {

	// 播放域名
	PlayDomain string `json:"play_domain"`

	// 租户ID
	ProjectId string `json:"project_id"`

	// 租户名
	TenantName string `json:"tenant_name"`

	// 证书开始时间，格式：YYYY-MM-DDTHH:mm:ssZ，UTC时间
	StartTime string `json:"start_time"`

	// 证书过期时间，格式：YYYY-MM-DDTHH:mm:ssZ，UTC时间
	ExpireTime string `json:"expire_time"`

	// 证书类型, 1：国密证书， 0:默认，标准国际证书
	CertType *int32 `json:"cert_type,omitempty"`

	// 国密证书类型, 1：签名证书， 2:加密证书
	GmCertType *int32 `json:"gm_cert_type,omitempty"`

	// 证书来源。 - scm：云证书与管理服务CCM。 - user：自有证书。
	Source *string `json:"source,omitempty"`

	// scm来源的证书名, 自有证书为空
	CertName *string `json:"cert_name,omitempty"`

	// scm来源的证书ID, 自有证书为空
	CertId *string `json:"cert_id,omitempty"`
}

func (o CertInfoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertInfoResp struct{}"
	}

	return strings.Join([]string{"CertInfoResp", string(data)}, " ")
}
