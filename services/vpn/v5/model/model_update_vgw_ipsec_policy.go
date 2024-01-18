package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateVgwIpsecPolicy struct {

	// 加密算法
	AuthenticationAlgorithm *UpdateVgwIpsecPolicyAuthenticationAlgorithm `json:"authentication_algorithm,omitempty"`

	// 加密算法
	EncryptionAlgorithm *UpdateVgwIpsecPolicyEncryptionAlgorithm `json:"encryption_algorithm,omitempty"`

	// PFS使用的DH密钥组
	Pfs *UpdateVgwIpsecPolicyPfs `json:"pfs,omitempty"`

	// 表示配置IPSec连接建立的隧道以时间为基准的生存周期
	LifetimeSeconds *int32 `json:"lifetime_seconds,omitempty"`
}

func (o UpdateVgwIpsecPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVgwIpsecPolicy struct{}"
	}

	return strings.Join([]string{"UpdateVgwIpsecPolicy", string(data)}, " ")
}

type UpdateVgwIpsecPolicyAuthenticationAlgorithm struct {
	value string
}

type UpdateVgwIpsecPolicyAuthenticationAlgorithmEnum struct {
	SHA2_256 UpdateVgwIpsecPolicyAuthenticationAlgorithm
	SHA2_384 UpdateVgwIpsecPolicyAuthenticationAlgorithm
	SHA2_512 UpdateVgwIpsecPolicyAuthenticationAlgorithm
}

func GetUpdateVgwIpsecPolicyAuthenticationAlgorithmEnum() UpdateVgwIpsecPolicyAuthenticationAlgorithmEnum {
	return UpdateVgwIpsecPolicyAuthenticationAlgorithmEnum{
		SHA2_256: UpdateVgwIpsecPolicyAuthenticationAlgorithm{
			value: "sha2-256",
		},
		SHA2_384: UpdateVgwIpsecPolicyAuthenticationAlgorithm{
			value: "sha2-384",
		},
		SHA2_512: UpdateVgwIpsecPolicyAuthenticationAlgorithm{
			value: "sha2-512",
		},
	}
}

func (c UpdateVgwIpsecPolicyAuthenticationAlgorithm) Value() string {
	return c.value
}

func (c UpdateVgwIpsecPolicyAuthenticationAlgorithm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateVgwIpsecPolicyAuthenticationAlgorithm) UnmarshalJSON(b []byte) error {
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

type UpdateVgwIpsecPolicyEncryptionAlgorithm struct {
	value string
}

type UpdateVgwIpsecPolicyEncryptionAlgorithmEnum struct {
	AES_128        UpdateVgwIpsecPolicyEncryptionAlgorithm
	AES_192        UpdateVgwIpsecPolicyEncryptionAlgorithm
	AES_256        UpdateVgwIpsecPolicyEncryptionAlgorithm
	AES_128_GCM_16 UpdateVgwIpsecPolicyEncryptionAlgorithm
	AES_256_GCM_16 UpdateVgwIpsecPolicyEncryptionAlgorithm
}

func GetUpdateVgwIpsecPolicyEncryptionAlgorithmEnum() UpdateVgwIpsecPolicyEncryptionAlgorithmEnum {
	return UpdateVgwIpsecPolicyEncryptionAlgorithmEnum{
		AES_128: UpdateVgwIpsecPolicyEncryptionAlgorithm{
			value: "aes-128",
		},
		AES_192: UpdateVgwIpsecPolicyEncryptionAlgorithm{
			value: "aes-192",
		},
		AES_256: UpdateVgwIpsecPolicyEncryptionAlgorithm{
			value: "aes-256",
		},
		AES_128_GCM_16: UpdateVgwIpsecPolicyEncryptionAlgorithm{
			value: "aes-128-gcm-16",
		},
		AES_256_GCM_16: UpdateVgwIpsecPolicyEncryptionAlgorithm{
			value: "aes-256-gcm-16",
		},
	}
}

func (c UpdateVgwIpsecPolicyEncryptionAlgorithm) Value() string {
	return c.value
}

func (c UpdateVgwIpsecPolicyEncryptionAlgorithm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateVgwIpsecPolicyEncryptionAlgorithm) UnmarshalJSON(b []byte) error {
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

type UpdateVgwIpsecPolicyPfs struct {
	value string
}

type UpdateVgwIpsecPolicyPfsEnum struct {
	GROUP14 UpdateVgwIpsecPolicyPfs
	GROUP15 UpdateVgwIpsecPolicyPfs
	GROUP16 UpdateVgwIpsecPolicyPfs
	GROUP19 UpdateVgwIpsecPolicyPfs
	GROUP20 UpdateVgwIpsecPolicyPfs
	GROUP21 UpdateVgwIpsecPolicyPfs
	DISABLE UpdateVgwIpsecPolicyPfs
}

func GetUpdateVgwIpsecPolicyPfsEnum() UpdateVgwIpsecPolicyPfsEnum {
	return UpdateVgwIpsecPolicyPfsEnum{
		GROUP14: UpdateVgwIpsecPolicyPfs{
			value: "group14",
		},
		GROUP15: UpdateVgwIpsecPolicyPfs{
			value: "group15",
		},
		GROUP16: UpdateVgwIpsecPolicyPfs{
			value: "group16",
		},
		GROUP19: UpdateVgwIpsecPolicyPfs{
			value: "group19",
		},
		GROUP20: UpdateVgwIpsecPolicyPfs{
			value: "group20",
		},
		GROUP21: UpdateVgwIpsecPolicyPfs{
			value: "group21",
		},
		DISABLE: UpdateVgwIpsecPolicyPfs{
			value: "disable",
		},
	}
}

func (c UpdateVgwIpsecPolicyPfs) Value() string {
	return c.value
}

func (c UpdateVgwIpsecPolicyPfs) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateVgwIpsecPolicyPfs) UnmarshalJSON(b []byte) error {
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
