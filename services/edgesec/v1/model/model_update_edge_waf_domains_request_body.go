package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateEdgeWafDomainsRequestBody 更新防护域名的请求
type UpdateEdgeWafDomainsRequestBody struct {

	// 防护状态
	ProtectStatus *int32 `json:"protect_status,omitempty"`

	// 接入状态
	AccessStatus *int32 `json:"access_status,omitempty"`

	// 域名名称
	WebTag *string `json:"web_tag,omitempty"`

	// 域名描述
	Description *string `json:"description,omitempty"`

	// 证书id，通过查询证书列表接口（ListCertificates）接口获取证书id   - 对外协议为HTTP时不需要填写   - 对外协议HTTPS时为必填参数   - 查询证书列表接口未开放时，从边缘安全控制台->边缘WAF->证书管理获取
	CertificateId *string `json:"certificate_id,omitempty"`

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 配置的最低TLS版本（TLS v1.0/TLS v1.1/TLS v1.2）,默认为TLS v1.0版本，对外协议为https时才有tls参数
	Tls *UpdateEdgeWafDomainsRequestBodyTls `json:"tls,omitempty"`

	// 对外协议为https时才有cipher参数，加密套件（cipher_1，cipher_2，cipher_3，cipher_4，cipher_default）：  - cipher_1： 加密算法为ECDHE-ECDSA-AES256-GCM-SHA384:HIGH:!MEDIUM:!LOW:!aNULL:!eNULL:!DES:!MD5:!PSK:!RC4:!kRSA:!SRP:!3DES:!DSS:!EXP:!CAMELLIA:@STRENGTH   - cipher_2：加密算法为EECDH+AESGCM:EDH+AESGCM   - cipher_3：加密算法为ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-SHA384:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!DH:!EDH   - cipher_4：加密算法为ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-SHA384:AES256-SHA256:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!EDH   - cipher_default： 加密算法为ECDHE-RSA-AES256-SHA384:AES256-SHA256:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!DH:!EDH:!AESGCM
	Cipher *UpdateEdgeWafDomainsRequestBodyCipher `json:"cipher,omitempty"`

	BlockPage *WafBlockPage `json:"block_page,omitempty"`

	TrafficMark *WafTrafficMark `json:"traffic_mark,omitempty"`

	Flag *Flag `json:"flag,omitempty"`

	// 域名可扩展字段
	Extend map[string]string `json:"extend,omitempty"`
}

func (o UpdateEdgeWafDomainsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEdgeWafDomainsRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateEdgeWafDomainsRequestBody", string(data)}, " ")
}

type UpdateEdgeWafDomainsRequestBodyTls struct {
	value string
}

type UpdateEdgeWafDomainsRequestBodyTlsEnum struct {
	TLS_V1_0 UpdateEdgeWafDomainsRequestBodyTls
	TLS_V1_1 UpdateEdgeWafDomainsRequestBodyTls
	TLS_V1_2 UpdateEdgeWafDomainsRequestBodyTls
}

func GetUpdateEdgeWafDomainsRequestBodyTlsEnum() UpdateEdgeWafDomainsRequestBodyTlsEnum {
	return UpdateEdgeWafDomainsRequestBodyTlsEnum{
		TLS_V1_0: UpdateEdgeWafDomainsRequestBodyTls{
			value: "TLS v1.0",
		},
		TLS_V1_1: UpdateEdgeWafDomainsRequestBodyTls{
			value: "TLS v1.1",
		},
		TLS_V1_2: UpdateEdgeWafDomainsRequestBodyTls{
			value: "TLS v1.2",
		},
	}
}

func (c UpdateEdgeWafDomainsRequestBodyTls) Value() string {
	return c.value
}

func (c UpdateEdgeWafDomainsRequestBodyTls) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateEdgeWafDomainsRequestBodyTls) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type UpdateEdgeWafDomainsRequestBodyCipher struct {
	value string
}

type UpdateEdgeWafDomainsRequestBodyCipherEnum struct {
	CIPHER_1       UpdateEdgeWafDomainsRequestBodyCipher
	CIPHER_2       UpdateEdgeWafDomainsRequestBodyCipher
	CIPHER_3       UpdateEdgeWafDomainsRequestBodyCipher
	CIPHER_4       UpdateEdgeWafDomainsRequestBodyCipher
	CIPHER_DEFAULT UpdateEdgeWafDomainsRequestBodyCipher
}

func GetUpdateEdgeWafDomainsRequestBodyCipherEnum() UpdateEdgeWafDomainsRequestBodyCipherEnum {
	return UpdateEdgeWafDomainsRequestBodyCipherEnum{
		CIPHER_1: UpdateEdgeWafDomainsRequestBodyCipher{
			value: "cipher_1",
		},
		CIPHER_2: UpdateEdgeWafDomainsRequestBodyCipher{
			value: "cipher_2",
		},
		CIPHER_3: UpdateEdgeWafDomainsRequestBodyCipher{
			value: "cipher_3",
		},
		CIPHER_4: UpdateEdgeWafDomainsRequestBodyCipher{
			value: "cipher_4",
		},
		CIPHER_DEFAULT: UpdateEdgeWafDomainsRequestBodyCipher{
			value: "cipher_default",
		},
	}
}

func (c UpdateEdgeWafDomainsRequestBodyCipher) Value() string {
	return c.value
}

func (c UpdateEdgeWafDomainsRequestBodyCipher) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateEdgeWafDomainsRequestBodyCipher) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
