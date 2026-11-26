package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HttpsDetailCertificates 双证书场景下的单个证书信息
type HttpsDetailCertificates struct {

	// **参数解释：** 证书类型 **约束限制：** 不涉及 **取值范围：** - server: 国际证书 - server_sm: 国密证书  **默认取值：** 不涉及
	CertificateType *string `json:"certificate_type,omitempty"`

	// **参数解释：** 证书来源 **约束限制：** 不涉及 **取值范围：** - 0: 自有证书 - 1: 云托管证书 - 2: SCM证书  **默认取值：** 不涉及
	CertificateSource *int32 `json:"certificate_source,omitempty"`

	// **参数解释：** 证书名称 **约束限制：** 不涉及 **取值范围：** 不涉及
	CertificateName *string `json:"certificate_name,omitempty"`

	// **参数解释：** 证书内容 **约束限制：** 不涉及 **取值范围：** 不涉及
	CertificateValue *string `json:"certificate_value,omitempty"`

	// **参数解释：** 私钥内容 **约束限制：** 不涉及 **取值范围：** 不涉及
	PrivateKey *string `json:"private_key,omitempty"`

	// **参数解释：** 国密证书加密证书内容 **约束限制：** 不涉及 **取值范围：** 不涉及
	EncCertificateValue *string `json:"enc_certificate_value,omitempty"`

	// **参数解释：** 国密证书加密私钥内容 **约束限制：** 不涉及 **取值范围：** 不涉及
	EncPrivateKey *string `json:"enc_private_key,omitempty"`

	// **参数解释：** 证书过期时间，Unix时间戳格式，单位为毫秒 **约束限制：** 不涉及 **取值范围：** 不涉及
	ExpireTime *int64 `json:"expire_time,omitempty"`

	// **参数解释：** SCM证书ID **约束限制：** 不涉及 **取值范围：** 不涉及
	ScmCertificateId *string `json:"scm_certificate_id,omitempty"`
}

func (o HttpsDetailCertificates) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpsDetailCertificates struct{}"
	}

	return strings.Join([]string{"HttpsDetailCertificates", string(data)}, " ")
}
