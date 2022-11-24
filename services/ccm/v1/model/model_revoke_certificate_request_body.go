package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RevokeCertificateRequestBody struct {

	// 吊销理由，支持以下选项：   - **UNSPECIFIED** : 吊销时未指定吊销原因，为默认值；   - **KEY_COMPROMISE** : 证书密钥材料泄露；   - **CERTIFICATE_AUTHORITY_COMPROMISE** : 颁发路径上，可能存在CA密钥材料泄露；   - **AFFILIATION_CHANGED** : 证书中的主体或其他信息已经被改变；   - **SUPERSEDED** : 此证书已被取代；   - **CESSATION_OF_OPERATION** : 此证书或颁发路径中的实体已停止运营；   - **CERTIFICATE_HOLD** : 该证书当前不应被视为有效，预计将来可能会生效；   - **PRIVILEGE_WITHDRAWN** : 此证书不再具备其声明的属性的权限；   - **ATTRIBUTE_AUTHORITY_COMPROMISE** : 担保此证书属性的机构可能已受到损害。 > 当不想填写吊销理由时，请求body体请置为\"{}\"，否则将会报异常，默认值为UNSPECIFIED。
	Reason *string `json:"reason,omitempty"`
}

func (o RevokeCertificateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RevokeCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"RevokeCertificateRequestBody", string(data)}, " ")
}
