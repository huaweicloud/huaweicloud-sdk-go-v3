package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MultiCertificateBatchVo **参数解释：** 双证书场景下的单个证书配置项 **约束限制：** 不涉及
type MultiCertificateBatchVo struct {

	// **参数解释：** 证书名称 **约束限制：** 不涉及 **取值范围：** - 3-64个字符 - 由英文字母、数字及“-”、“_”、空格组成  **默认取值：** 不涉及
	CertName *string `json:"cert_name,omitempty"`

	// **参数解释：** 证书类型 **约束限制：** 不涉及 **取值范围：** - 0: 国际证书 - 1: 国密证书  **默认取值：** 不涉及
	CertificateCategory *int32 `json:"certificate_category,omitempty"`

	// **参数解释：** 证书类型 **约束限制：** 不涉及 **取值范围：** - 0: 自有证书 - 1: 云托管证书 - 2: SCM证书  **默认取值：** 不涉及
	CertificateType *int32 `json:"certificate_type,omitempty"`

	// **参数解释：** SSL证书内容 **约束限制：** - 仅支持PEM编码格式 - 需要传入完整证书链  **取值范围：** 证书链长度不超过20KB **默认取值：** 不涉及
	Certificate *string `json:"certificate,omitempty"`

	// **参数解释：** SSL证书私钥内容 **约束限制：** - 仅支持PEM编码格式  **取值范围：** 不涉及 **默认取值：** 不涉及
	PrivateKey *string `json:"private_key,omitempty"`

	// **参数解释：** 国密证书加密证书内容 **约束限制：** certificate_category 为 1 且 certificate_type 为 0 时必传，须与 enc_private_key 同时传入 **取值范围：** 不涉及 **默认取值：** 不涉及
	EncCertificate *string `json:"enc_certificate,omitempty"`

	// **参数解释：** 国密证书加密私钥内容 **约束限制：** certificate_category 为 1 且 certificate_type 为 0 时必传，须与 enc_certificate 同时传入 **取值范围：** 不涉及 **默认取值：** 不涉及
	EncPrivateKey *string `json:"enc_private_key,omitempty"`

	// **参数解释：** SCM证书ID，配置SCM证书时必传 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ScmCertificateId *string `json:"scm_certificate_id,omitempty"`
}

func (o MultiCertificateBatchVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiCertificateBatchVo struct{}"
	}

	return strings.Join([]string{"MultiCertificateBatchVo", string(data)}, " ")
}
