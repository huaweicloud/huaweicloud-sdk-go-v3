package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClusterCert struct {

	// **参数解释**： 上下文user信息 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Server *string `json:"server,omitempty"`

	// **参数解释**： 证书授权数据 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	CertificateAuthorityData *string `json:"certificate-authority-data,omitempty"`

	// **参数解释**： 不校验服务端证书 **约束限制**： 不涉及 **取值范围**： - true：跳过校验服务端证书 - false：校验服务端证书  **默认取值**： 在 cluster 类型为 externalCluster 时，该值为 true。
	InsecureSkipTlsVerify *bool `json:"insecure-skip-tls-verify,omitempty"`
}

func (o ClusterCert) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterCert struct{}"
	}

	return strings.Join([]string{"ClusterCert", string(data)}, " ")
}
