package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachOrDetachDomainInfo 证书绑定或解绑域名信息。如果填了instance_id则只操作该实例下指定域名；如果不填instance_id则操作全局指定域名。
type AttachOrDetachDomainInfo struct {

	// 域名
	Domain string `json:"domain"`

	// 实例ID集合
	InstanceIds *[]string `json:"instance_ids,omitempty"`

	// 是否开启客户端证书校验。当绑定证书存在trusted_root_ca时，默认开启；当绑定证书不存在trusted_root_ca时，默认关闭。
	VerifiedClientCertificateEnabled *bool `json:"verified_client_certificate_enabled,omitempty"`
}

func (o AttachOrDetachDomainInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachOrDetachDomainInfo struct{}"
	}

	return strings.Join([]string{"AttachOrDetachDomainInfo", string(data)}, " ")
}
