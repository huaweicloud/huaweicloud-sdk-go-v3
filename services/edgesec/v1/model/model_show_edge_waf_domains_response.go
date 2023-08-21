package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowEdgeWafDomainsResponse Response Object
type ShowEdgeWafDomainsResponse struct {

	// 域名id
	Id *string `json:"id,omitempty"`

	// 域名
	DomainName *string `json:"domain_name,omitempty"`

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 租户ID
	TenantId *string `json:"tenant_id,omitempty"`

	// 上一次开启防护的时间
	OpenTime *int64 `json:"open_time,omitempty"`

	// 上一次关闭防护的时间
	CloseTime *int64 `json:"close_time,omitempty"`

	// cdn域名调度情况（0：未防护，1：配置中，2：已防护，3：删除中）
	DispatchStatus *int32 `json:"dispatch_status,omitempty"`

	// 域名在CDN所属区域
	ServiceArea *string `json:"service_area,omitempty"`

	// 域名名称
	WebTag *string `json:"web_tag,omitempty"`

	// 域名描述
	Description *string `json:"description,omitempty"`

	// 策略id
	PolicyId *string `json:"policy_id,omitempty"`

	// 协议
	Protocol *ShowEdgeWafDomainsResponseProtocol `json:"protocol,omitempty"`

	// 证书id
	CertificateId *string `json:"certificate_id,omitempty"`

	// 证书名称
	CertificateName *string `json:"certificate_name,omitempty"`

	// 配置的最低TLS版本（TLS v1.0/TLS v1.1/TLS v1.2），默认为TLS v1.0版本，对外协议为https时才有tls参数
	Tls *ShowEdgeWafDomainsResponseTls `json:"tls,omitempty"`

	// 对外协议为https时才有cipher参数，加密套件（cipher_1，cipher_2，cipher_3，cipher_4，cipher_default）：  - cipher_1： 加密算法为ECDHE-ECDSA-AES256-GCM-SHA384:HIGH:!MEDIUM:!LOW:!aNULL:!eNULL:!DES:!MD5:!PSK:!RC4:!kRSA:!SRP:!3DES:!DSS:!EXP:!CAMELLIA:@STRENGTH   - cipher_2：加密算法为EECDH+AESGCM:EDH+AESGCM   - cipher_3：加密算法为ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-SHA384:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!DH:!EDH   - cipher_4：加密算法为ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-SHA384:AES256-SHA256:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!EDH   - cipher_default： 加密算法为ECDHE-RSA-AES256-SHA384:AES256-SHA256:RC4:HIGH:!MD5:!aNULL:!eNULL:!NULL:!DH:!EDH:!AESGCM
	Cipher *ShowEdgeWafDomainsResponseCipher `json:"cipher,omitempty"`

	// 防护状态：  - 0-关闭   - 1-开启
	ProtectStatus *int32 `json:"protect_status,omitempty"`

	// 接入状态：  - 0-未接入   - 1-已接入
	AccessStatus *int32 `json:"access_status,omitempty"`

	// 创建域名的时间，13位时间戳
	CreateTime *int64 `json:"create_time,omitempty"`

	BlockPage *WafBlockPage `json:"block_page,omitempty"`

	TrafficMark *WafTrafficMark `json:"traffic_mark,omitempty"`

	Flag *Flag `json:"flag,omitempty"`

	// 域名可扩展属性
	Extend map[string]string `json:"extend,omitempty"`

	// 是否为ddos防护域名
	IsAdded        *bool `json:"is_added,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowEdgeWafDomainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEdgeWafDomainsResponse struct{}"
	}

	return strings.Join([]string{"ShowEdgeWafDomainsResponse", string(data)}, " ")
}

type ShowEdgeWafDomainsResponseProtocol struct {
	value string
}

type ShowEdgeWafDomainsResponseProtocolEnum struct {
	HTTP  ShowEdgeWafDomainsResponseProtocol
	HTTPS ShowEdgeWafDomainsResponseProtocol
}

func GetShowEdgeWafDomainsResponseProtocolEnum() ShowEdgeWafDomainsResponseProtocolEnum {
	return ShowEdgeWafDomainsResponseProtocolEnum{
		HTTP: ShowEdgeWafDomainsResponseProtocol{
			value: "http",
		},
		HTTPS: ShowEdgeWafDomainsResponseProtocol{
			value: "https",
		},
	}
}

func (c ShowEdgeWafDomainsResponseProtocol) Value() string {
	return c.value
}

func (c ShowEdgeWafDomainsResponseProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEdgeWafDomainsResponseProtocol) UnmarshalJSON(b []byte) error {
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

type ShowEdgeWafDomainsResponseTls struct {
	value string
}

type ShowEdgeWafDomainsResponseTlsEnum struct {
	TLS_V1_0 ShowEdgeWafDomainsResponseTls
	TLS_V1_1 ShowEdgeWafDomainsResponseTls
	TLS_V1_2 ShowEdgeWafDomainsResponseTls
}

func GetShowEdgeWafDomainsResponseTlsEnum() ShowEdgeWafDomainsResponseTlsEnum {
	return ShowEdgeWafDomainsResponseTlsEnum{
		TLS_V1_0: ShowEdgeWafDomainsResponseTls{
			value: "TLS v1.0",
		},
		TLS_V1_1: ShowEdgeWafDomainsResponseTls{
			value: "TLS v1.1",
		},
		TLS_V1_2: ShowEdgeWafDomainsResponseTls{
			value: "TLS v1.2",
		},
	}
}

func (c ShowEdgeWafDomainsResponseTls) Value() string {
	return c.value
}

func (c ShowEdgeWafDomainsResponseTls) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEdgeWafDomainsResponseTls) UnmarshalJSON(b []byte) error {
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

type ShowEdgeWafDomainsResponseCipher struct {
	value string
}

type ShowEdgeWafDomainsResponseCipherEnum struct {
	CIPHER_1       ShowEdgeWafDomainsResponseCipher
	CIPHER_2       ShowEdgeWafDomainsResponseCipher
	CIPHER_3       ShowEdgeWafDomainsResponseCipher
	CIPHER_4       ShowEdgeWafDomainsResponseCipher
	CIPHER_DEFAULT ShowEdgeWafDomainsResponseCipher
}

func GetShowEdgeWafDomainsResponseCipherEnum() ShowEdgeWafDomainsResponseCipherEnum {
	return ShowEdgeWafDomainsResponseCipherEnum{
		CIPHER_1: ShowEdgeWafDomainsResponseCipher{
			value: "cipher_1",
		},
		CIPHER_2: ShowEdgeWafDomainsResponseCipher{
			value: "cipher_2",
		},
		CIPHER_3: ShowEdgeWafDomainsResponseCipher{
			value: "cipher_3",
		},
		CIPHER_4: ShowEdgeWafDomainsResponseCipher{
			value: "cipher_4",
		},
		CIPHER_DEFAULT: ShowEdgeWafDomainsResponseCipher{
			value: "cipher_default",
		},
	}
}

func (c ShowEdgeWafDomainsResponseCipher) Value() string {
	return c.value
}

func (c ShowEdgeWafDomainsResponseCipher) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEdgeWafDomainsResponseCipher) UnmarshalJSON(b []byte) error {
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
