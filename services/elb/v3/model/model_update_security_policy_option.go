package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 更新自定义安全策略的请求参数。
type UpdateSecurityPolicyOption struct {

	// 自定义安全策略的名称。
	Name *string `json:"name,omitempty"`

	// 自定义安全策略的描述信息。
	Description *string `json:"description,omitempty"`

	// 自定义安全策略选择的TLS协议列表。取值：TLSv1, TLSv1.1, TLSv1.2, TLSv1.3
	Protocols *[]string `json:"protocols,omitempty"`

	// 自定义安全策略的加密套件列表。支持以下加密套件：   ECDHE-RSA-AES256-GCM-SHA384,ECDHE-RSA-AES128-GCM-SHA256,ECDHE-ECDSA-AES256-GCM-SHA384,ECDHE-ECDSA-AES128-GCM-SHA256,AES128-GCM-SHA256,AES256-GCM-SHA384,ECDHE-ECDSA-AES128-SHA256,ECDHE-RSA-AES128-SHA256,AES128-SHA256,AES256-SHA256,ECDHE-ECDSA-AES256-SHA384,ECDHE-RSA-AES256-SHA384,ECDHE-ECDSA-AES128-SHA,ECDHE-RSA-AES128-SHA,ECDHE-RSA-AES256-SHA,ECDHE-ECDSA-AES256-SHA,AES128-SHA,AES256-SHA,CAMELLIA128-SHA,DES-CBC3-SHA,CAMELLIA256-SHA,ECDHE-RSA-CHACHA20-POLY1305,ECDHE-ECDSA-CHACHA20-POLY1305,TLS_AES_128_GCM_SHA256,TLS_AES_256_GCM_SHA384,TLS_CHACHA20_POLY1305_SHA256,TLS_AES_128_CCM_SHA256,TLS_AES_128_CCM_8_SHA256   使用说明：   - 协议和加密套件必须匹配，即ciphers中必须至少有一种有与协议匹配的加密套件。   > 协议与加密套件的匹配关系可参考系统安全策略
	Ciphers *[]UpdateSecurityPolicyOptionCiphers `json:"ciphers,omitempty"`
}

func (o UpdateSecurityPolicyOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityPolicyOption struct{}"
	}

	return strings.Join([]string{"UpdateSecurityPolicyOption", string(data)}, " ")
}

type UpdateSecurityPolicyOptionCiphers struct {
	value string
}

type UpdateSecurityPolicyOptionCiphersEnum struct {
	ECDHE_RSA_AES256_GCM_SHA384   UpdateSecurityPolicyOptionCiphers
	ECDHE_RSA_AES128_GCM_SHA256   UpdateSecurityPolicyOptionCiphers
	ECDHE_ECDSA_AES256_GCM_SHA384 UpdateSecurityPolicyOptionCiphers
	ECDHE_ECDSA_AES128_GCM_SHA256 UpdateSecurityPolicyOptionCiphers
	AES128_GCM_SHA256             UpdateSecurityPolicyOptionCiphers
	AES256_GCM_SHA384             UpdateSecurityPolicyOptionCiphers
	ECDHE_ECDSA_AES128_SHA256     UpdateSecurityPolicyOptionCiphers
	ECDHE_RSA_AES128_SHA256       UpdateSecurityPolicyOptionCiphers
	AES128_SHA256                 UpdateSecurityPolicyOptionCiphers
	AES256_SHA256                 UpdateSecurityPolicyOptionCiphers
	ECDHE_ECDSA_AES256_SHA384     UpdateSecurityPolicyOptionCiphers
	ECDHE_RSA_AES256_SHA384       UpdateSecurityPolicyOptionCiphers
	ECDHE_ECDSA_AES128_SHA        UpdateSecurityPolicyOptionCiphers
	ECDHE_RSA_AES128_SHA          UpdateSecurityPolicyOptionCiphers
	ECDHE_RSA_AES256_SHA          UpdateSecurityPolicyOptionCiphers
	ECDHE_ECDSA_AES256_SHA        UpdateSecurityPolicyOptionCiphers
	AES128_SHA                    UpdateSecurityPolicyOptionCiphers
	AES256_SHA                    UpdateSecurityPolicyOptionCiphers
	CAMELLIA128_SHA               UpdateSecurityPolicyOptionCiphers
	DES_CBC3_SHA                  UpdateSecurityPolicyOptionCiphers
	CAMELLIA256_SHA               UpdateSecurityPolicyOptionCiphers
	ECDHE_RSA_CHACHA20_POLY1305   UpdateSecurityPolicyOptionCiphers
	ECDHE_ECDSA_CHACHA20_POLY1305 UpdateSecurityPolicyOptionCiphers
	TLS_AES_128_GCM_SHA256        UpdateSecurityPolicyOptionCiphers
	TLS_AES_256_GCM_SHA384        UpdateSecurityPolicyOptionCiphers
	TLS_CHACHA20_POLY1305_SHA256  UpdateSecurityPolicyOptionCiphers
	TLS_AES_128_CCM_SHA256        UpdateSecurityPolicyOptionCiphers
	TLS_AES_128_CCM_8_SHA256      UpdateSecurityPolicyOptionCiphers
	ECC_SM4_SM3                   UpdateSecurityPolicyOptionCiphers
	ECDHE_SM4_SM3                 UpdateSecurityPolicyOptionCiphers
}

func GetUpdateSecurityPolicyOptionCiphersEnum() UpdateSecurityPolicyOptionCiphersEnum {
	return UpdateSecurityPolicyOptionCiphersEnum{
		ECDHE_RSA_AES256_GCM_SHA384: UpdateSecurityPolicyOptionCiphers{
			value: "ECDHE-RSA-AES256-GCM-SHA384",
		},
		ECDHE_RSA_AES128_GCM_SHA256: UpdateSecurityPolicyOptionCiphers{
			value: "ECDHE-RSA-AES128-GCM-SHA256",
		},
		ECDHE_ECDSA_AES256_GCM_SHA384: UpdateSecurityPolicyOptionCiphers{
			value: "ECDHE-ECDSA-AES256-GCM-SHA384",
		},
		ECDHE_ECDSA_AES128_GCM_SHA256: UpdateSecurityPolicyOptionCiphers{
			value: "ECDHE-ECDSA-AES128-GCM-SHA256",
		},
		AES128_GCM_SHA256: UpdateSecurityPolicyOptionCiphers{
			value: "AES128-GCM-SHA256",
		},
		AES256_GCM_SHA384: UpdateSecurityPolicyOptionCiphers{
			value: "AES256-GCM-SHA384",
		},
		ECDHE_ECDSA_AES128_SHA256: UpdateSecurityPolicyOptionCiphers{
			value: "ECDHE-ECDSA-AES128-SHA256",
		},
		ECDHE_RSA_AES128_SHA256: UpdateSecurityPolicyOptionCiphers{
			value: "ECDHE-RSA-AES128-SHA256",
		},
		AES128_SHA256: UpdateSecurityPolicyOptionCiphers{
			value: "AES128-SHA256",
		},
		AES256_SHA256: UpdateSecurityPolicyOptionCiphers{
			value: "AES256-SHA256",
		},
		ECDHE_ECDSA_AES256_SHA384: UpdateSecurityPolicyOptionCiphers{
			value: "ECDHE-ECDSA-AES256-SHA384",
		},
		ECDHE_RSA_AES256_SHA384: UpdateSecurityPolicyOptionCiphers{
			value: "ECDHE-RSA-AES256-SHA384",
		},
		ECDHE_ECDSA_AES128_SHA: UpdateSecurityPolicyOptionCiphers{
			value: "ECDHE-ECDSA-AES128-SHA",
		},
		ECDHE_RSA_AES128_SHA: UpdateSecurityPolicyOptionCiphers{
			value: "ECDHE-RSA-AES128-SHA",
		},
		ECDHE_RSA_AES256_SHA: UpdateSecurityPolicyOptionCiphers{
			value: "ECDHE-RSA-AES256-SHA",
		},
		ECDHE_ECDSA_AES256_SHA: UpdateSecurityPolicyOptionCiphers{
			value: "ECDHE-ECDSA-AES256-SHA",
		},
		AES128_SHA: UpdateSecurityPolicyOptionCiphers{
			value: "AES128-SHA",
		},
		AES256_SHA: UpdateSecurityPolicyOptionCiphers{
			value: "AES256-SHA",
		},
		CAMELLIA128_SHA: UpdateSecurityPolicyOptionCiphers{
			value: "CAMELLIA128-SHA",
		},
		DES_CBC3_SHA: UpdateSecurityPolicyOptionCiphers{
			value: "DES-CBC3-SHA",
		},
		CAMELLIA256_SHA: UpdateSecurityPolicyOptionCiphers{
			value: "CAMELLIA256-SHA",
		},
		ECDHE_RSA_CHACHA20_POLY1305: UpdateSecurityPolicyOptionCiphers{
			value: "ECDHE-RSA-CHACHA20-POLY1305",
		},
		ECDHE_ECDSA_CHACHA20_POLY1305: UpdateSecurityPolicyOptionCiphers{
			value: "ECDHE-ECDSA-CHACHA20-POLY1305",
		},
		TLS_AES_128_GCM_SHA256: UpdateSecurityPolicyOptionCiphers{
			value: "TLS_AES_128_GCM_SHA256",
		},
		TLS_AES_256_GCM_SHA384: UpdateSecurityPolicyOptionCiphers{
			value: "TLS_AES_256_GCM_SHA384",
		},
		TLS_CHACHA20_POLY1305_SHA256: UpdateSecurityPolicyOptionCiphers{
			value: "TLS_CHACHA20_POLY1305_SHA256",
		},
		TLS_AES_128_CCM_SHA256: UpdateSecurityPolicyOptionCiphers{
			value: "TLS_AES_128_CCM_SHA256",
		},
		TLS_AES_128_CCM_8_SHA256: UpdateSecurityPolicyOptionCiphers{
			value: "TLS_AES_128_CCM_8_SHA256",
		},
		ECC_SM4_SM3: UpdateSecurityPolicyOptionCiphers{
			value: "ECC-SM4-SM3",
		},
		ECDHE_SM4_SM3: UpdateSecurityPolicyOptionCiphers{
			value: "ECDHE-SM4-SM3",
		},
	}
}

func (c UpdateSecurityPolicyOptionCiphers) Value() string {
	return c.value
}

func (c UpdateSecurityPolicyOptionCiphers) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSecurityPolicyOptionCiphers) UnmarshalJSON(b []byte) error {
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
