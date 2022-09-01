package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 自定义安全策略创建参数。
type CreateSecurityPolicyOption struct {

	// 自定义安全策略的名称。默认空字符串\"\"。
	Name *string `json:"name,omitempty" xml:"name"`

	// 自定义安全策略的描述信息。默认空字符串\"\"。
	Description *string `json:"description,omitempty" xml:"description"`

	// 所属企业项目ID。  [不支持该字段，请勿使用。](tag:dt,dt_test,hcso_dt)
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	// 自定义安全策略选择的TLS协议列表。取值：TLSv1, TLSv1.1, TLSv1.2, TLSv1.3
	Protocols []string `json:"protocols" xml:"protocols"`

	// 自定义策略的加密套件列表。支持以下加密套件：   ECDHE-RSA-AES256-GCM-SHA384,ECDHE-RSA-AES128-GCM-SHA256,ECDHE-ECDSA-AES256-GCM-SHA384,ECDHE-ECDSA-AES128-GCM-SHA256,AES128-GCM-SHA256,AES256-GCM-SHA384,ECDHE-ECDSA-AES128-SHA256,ECDHE-RSA-AES128-SHA256,AES128-SHA256,AES256-SHA256,ECDHE-ECDSA-AES256-SHA384,ECDHE-RSA-AES256-SHA384,ECDHE-ECDSA-AES128-SHA,ECDHE-RSA-AES128-SHA,ECDHE-RSA-AES256-SHA,ECDHE-ECDSA-AES256-SHA,AES128-SHA,AES256-SHA,CAMELLIA128-SHA,DES-CBC3-SHA,CAMELLIA256-SHA,ECDHE-RSA-CHACHA20-POLY1305,ECDHE-ECDSA-CHACHA20-POLY1305,TLS_AES_128_GCM_SHA256,TLS_AES_256_GCM_SHA384,TLS_CHACHA20_POLY1305_SHA256,TLS_AES_128_CCM_SHA256,TLS_AES_128_CCM_8_SHA256    使用说明：   - 协议和加密套件必须匹配，即ciphers中必须至少有一种有与协议匹配的加密套件。   > 协议与加密套件的匹配关系可参考系统安全策略
	Ciphers []CreateSecurityPolicyOptionCiphers `json:"ciphers" xml:"ciphers"`
}

func (o CreateSecurityPolicyOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityPolicyOption struct{}"
	}

	return strings.Join([]string{"CreateSecurityPolicyOption", string(data)}, " ")
}

type CreateSecurityPolicyOptionCiphers struct {
	value string
}

type CreateSecurityPolicyOptionCiphersEnum struct {
	ECDHE_RSA_AES256_GCM_SHA384   CreateSecurityPolicyOptionCiphers
	ECDHE_RSA_AES128_GCM_SHA256   CreateSecurityPolicyOptionCiphers
	ECDHE_ECDSA_AES256_GCM_SHA384 CreateSecurityPolicyOptionCiphers
	ECDHE_ECDSA_AES128_GCM_SHA256 CreateSecurityPolicyOptionCiphers
	AES128_GCM_SHA256             CreateSecurityPolicyOptionCiphers
	AES256_GCM_SHA384             CreateSecurityPolicyOptionCiphers
	ECDHE_ECDSA_AES128_SHA256     CreateSecurityPolicyOptionCiphers
	ECDHE_RSA_AES128_SHA256       CreateSecurityPolicyOptionCiphers
	AES128_SHA256                 CreateSecurityPolicyOptionCiphers
	AES256_SHA256                 CreateSecurityPolicyOptionCiphers
	ECDHE_ECDSA_AES256_SHA384     CreateSecurityPolicyOptionCiphers
	ECDHE_RSA_AES256_SHA384       CreateSecurityPolicyOptionCiphers
	ECDHE_ECDSA_AES128_SHA        CreateSecurityPolicyOptionCiphers
	ECDHE_RSA_AES128_SHA          CreateSecurityPolicyOptionCiphers
	ECDHE_RSA_AES256_SHA          CreateSecurityPolicyOptionCiphers
	ECDHE_ECDSA_AES256_SHA        CreateSecurityPolicyOptionCiphers
	AES128_SHA                    CreateSecurityPolicyOptionCiphers
	AES256_SHA                    CreateSecurityPolicyOptionCiphers
	CAMELLIA128_SHA               CreateSecurityPolicyOptionCiphers
	DES_CBC3_SHA                  CreateSecurityPolicyOptionCiphers
	CAMELLIA256_SHA               CreateSecurityPolicyOptionCiphers
	ECDHE_RSA_CHACHA20_POLY1305   CreateSecurityPolicyOptionCiphers
	ECDHE_ECDSA_CHACHA20_POLY1305 CreateSecurityPolicyOptionCiphers
	TLS_AES_128_GCM_SHA256        CreateSecurityPolicyOptionCiphers
	TLS_AES_256_GCM_SHA384        CreateSecurityPolicyOptionCiphers
	TLS_CHACHA20_POLY1305_SHA256  CreateSecurityPolicyOptionCiphers
	TLS_AES_128_CCM_SHA256        CreateSecurityPolicyOptionCiphers
	TLS_AES_128_CCM_8_SHA256      CreateSecurityPolicyOptionCiphers
}

func GetCreateSecurityPolicyOptionCiphersEnum() CreateSecurityPolicyOptionCiphersEnum {
	return CreateSecurityPolicyOptionCiphersEnum{
		ECDHE_RSA_AES256_GCM_SHA384: CreateSecurityPolicyOptionCiphers{
			value: "ECDHE-RSA-AES256-GCM-SHA384",
		},
		ECDHE_RSA_AES128_GCM_SHA256: CreateSecurityPolicyOptionCiphers{
			value: "ECDHE-RSA-AES128-GCM-SHA256",
		},
		ECDHE_ECDSA_AES256_GCM_SHA384: CreateSecurityPolicyOptionCiphers{
			value: "ECDHE-ECDSA-AES256-GCM-SHA384",
		},
		ECDHE_ECDSA_AES128_GCM_SHA256: CreateSecurityPolicyOptionCiphers{
			value: "ECDHE-ECDSA-AES128-GCM-SHA256",
		},
		AES128_GCM_SHA256: CreateSecurityPolicyOptionCiphers{
			value: "AES128-GCM-SHA256",
		},
		AES256_GCM_SHA384: CreateSecurityPolicyOptionCiphers{
			value: "AES256-GCM-SHA384",
		},
		ECDHE_ECDSA_AES128_SHA256: CreateSecurityPolicyOptionCiphers{
			value: "ECDHE-ECDSA-AES128-SHA256",
		},
		ECDHE_RSA_AES128_SHA256: CreateSecurityPolicyOptionCiphers{
			value: "ECDHE-RSA-AES128-SHA256",
		},
		AES128_SHA256: CreateSecurityPolicyOptionCiphers{
			value: "AES128-SHA256",
		},
		AES256_SHA256: CreateSecurityPolicyOptionCiphers{
			value: "AES256-SHA256",
		},
		ECDHE_ECDSA_AES256_SHA384: CreateSecurityPolicyOptionCiphers{
			value: "ECDHE-ECDSA-AES256-SHA384",
		},
		ECDHE_RSA_AES256_SHA384: CreateSecurityPolicyOptionCiphers{
			value: "ECDHE-RSA-AES256-SHA384",
		},
		ECDHE_ECDSA_AES128_SHA: CreateSecurityPolicyOptionCiphers{
			value: "ECDHE-ECDSA-AES128-SHA",
		},
		ECDHE_RSA_AES128_SHA: CreateSecurityPolicyOptionCiphers{
			value: "ECDHE-RSA-AES128-SHA",
		},
		ECDHE_RSA_AES256_SHA: CreateSecurityPolicyOptionCiphers{
			value: "ECDHE-RSA-AES256-SHA",
		},
		ECDHE_ECDSA_AES256_SHA: CreateSecurityPolicyOptionCiphers{
			value: "ECDHE-ECDSA-AES256-SHA",
		},
		AES128_SHA: CreateSecurityPolicyOptionCiphers{
			value: "AES128-SHA",
		},
		AES256_SHA: CreateSecurityPolicyOptionCiphers{
			value: "AES256-SHA",
		},
		CAMELLIA128_SHA: CreateSecurityPolicyOptionCiphers{
			value: "CAMELLIA128-SHA",
		},
		DES_CBC3_SHA: CreateSecurityPolicyOptionCiphers{
			value: "DES-CBC3-SHA",
		},
		CAMELLIA256_SHA: CreateSecurityPolicyOptionCiphers{
			value: "CAMELLIA256-SHA",
		},
		ECDHE_RSA_CHACHA20_POLY1305: CreateSecurityPolicyOptionCiphers{
			value: "ECDHE-RSA-CHACHA20-POLY1305",
		},
		ECDHE_ECDSA_CHACHA20_POLY1305: CreateSecurityPolicyOptionCiphers{
			value: "ECDHE-ECDSA-CHACHA20-POLY1305",
		},
		TLS_AES_128_GCM_SHA256: CreateSecurityPolicyOptionCiphers{
			value: "TLS_AES_128_GCM_SHA256",
		},
		TLS_AES_256_GCM_SHA384: CreateSecurityPolicyOptionCiphers{
			value: "TLS_AES_256_GCM_SHA384",
		},
		TLS_CHACHA20_POLY1305_SHA256: CreateSecurityPolicyOptionCiphers{
			value: "TLS_CHACHA20_POLY1305_SHA256",
		},
		TLS_AES_128_CCM_SHA256: CreateSecurityPolicyOptionCiphers{
			value: "TLS_AES_128_CCM_SHA256",
		},
		TLS_AES_128_CCM_8_SHA256: CreateSecurityPolicyOptionCiphers{
			value: "TLS_AES_128_CCM_8_SHA256",
		},
	}
}

func (c CreateSecurityPolicyOptionCiphers) Value() string {
	return c.value
}

func (c CreateSecurityPolicyOptionCiphers) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSecurityPolicyOptionCiphers) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
