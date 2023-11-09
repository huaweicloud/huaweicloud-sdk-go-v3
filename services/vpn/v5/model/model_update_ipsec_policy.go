package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateIpsecPolicy struct {

	// 认证算法，SHA1和MD5安全性较低，请慎用
	AuthenticationAlgorithm *UpdateIpsecPolicyAuthenticationAlgorithm `json:"authentication_algorithm,omitempty"`

	// 加密算法，3DES安全性较低，请慎用
	EncryptionAlgorithm *UpdateIpsecPolicyEncryptionAlgorithm `json:"encryption_algorithm,omitempty"`

	// PFS使用的DH密钥组
	Pfs *string `json:"pfs,omitempty"`

	// 传输模式
	TransformProtocol *UpdateIpsecPolicyTransformProtocol `json:"transform_protocol,omitempty"`

	// 表示配置IPSec连接建立的隧道以时间为基准的生存周期
	LifetimeSeconds *int32 `json:"lifetime_seconds,omitempty"`

	// 封装模式，当前只有tunnel模式
	EncapsulationMode *UpdateIpsecPolicyEncapsulationMode `json:"encapsulation_mode,omitempty"`
}

func (o UpdateIpsecPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpsecPolicy struct{}"
	}

	return strings.Join([]string{"UpdateIpsecPolicy", string(data)}, " ")
}

type UpdateIpsecPolicyAuthenticationAlgorithm struct {
	value string
}

type UpdateIpsecPolicyAuthenticationAlgorithmEnum struct {
	SHA1     UpdateIpsecPolicyAuthenticationAlgorithm
	MD5      UpdateIpsecPolicyAuthenticationAlgorithm
	SHA2_256 UpdateIpsecPolicyAuthenticationAlgorithm
	SHA2_384 UpdateIpsecPolicyAuthenticationAlgorithm
	SHA2_512 UpdateIpsecPolicyAuthenticationAlgorithm
}

func GetUpdateIpsecPolicyAuthenticationAlgorithmEnum() UpdateIpsecPolicyAuthenticationAlgorithmEnum {
	return UpdateIpsecPolicyAuthenticationAlgorithmEnum{
		SHA1: UpdateIpsecPolicyAuthenticationAlgorithm{
			value: "sha1",
		},
		MD5: UpdateIpsecPolicyAuthenticationAlgorithm{
			value: "md5",
		},
		SHA2_256: UpdateIpsecPolicyAuthenticationAlgorithm{
			value: "sha2-256",
		},
		SHA2_384: UpdateIpsecPolicyAuthenticationAlgorithm{
			value: "sha2-384",
		},
		SHA2_512: UpdateIpsecPolicyAuthenticationAlgorithm{
			value: "sha2-512",
		},
	}
}

func (c UpdateIpsecPolicyAuthenticationAlgorithm) Value() string {
	return c.value
}

func (c UpdateIpsecPolicyAuthenticationAlgorithm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateIpsecPolicyAuthenticationAlgorithm) UnmarshalJSON(b []byte) error {
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

type UpdateIpsecPolicyEncryptionAlgorithm struct {
	value string
}

type UpdateIpsecPolicyEncryptionAlgorithmEnum struct {
	E_3DES          UpdateIpsecPolicyEncryptionAlgorithm
	AES_128         UpdateIpsecPolicyEncryptionAlgorithm
	AES_192         UpdateIpsecPolicyEncryptionAlgorithm
	AES_256         UpdateIpsecPolicyEncryptionAlgorithm
	AES_128_GCM_16  UpdateIpsecPolicyEncryptionAlgorithm
	AES_256_GCM_16  UpdateIpsecPolicyEncryptionAlgorithm
	AES_128_GCM_128 UpdateIpsecPolicyEncryptionAlgorithm
	AES_256_GCM_128 UpdateIpsecPolicyEncryptionAlgorithm
}

func GetUpdateIpsecPolicyEncryptionAlgorithmEnum() UpdateIpsecPolicyEncryptionAlgorithmEnum {
	return UpdateIpsecPolicyEncryptionAlgorithmEnum{
		E_3DES: UpdateIpsecPolicyEncryptionAlgorithm{
			value: "3des",
		},
		AES_128: UpdateIpsecPolicyEncryptionAlgorithm{
			value: "aes-128",
		},
		AES_192: UpdateIpsecPolicyEncryptionAlgorithm{
			value: "aes-192",
		},
		AES_256: UpdateIpsecPolicyEncryptionAlgorithm{
			value: "aes-256",
		},
		AES_128_GCM_16: UpdateIpsecPolicyEncryptionAlgorithm{
			value: "aes-128-gcm-16",
		},
		AES_256_GCM_16: UpdateIpsecPolicyEncryptionAlgorithm{
			value: "aes-256-gcm-16",
		},
		AES_128_GCM_128: UpdateIpsecPolicyEncryptionAlgorithm{
			value: "aes-128-gcm-128",
		},
		AES_256_GCM_128: UpdateIpsecPolicyEncryptionAlgorithm{
			value: "aes-256-gcm-128",
		},
	}
}

func (c UpdateIpsecPolicyEncryptionAlgorithm) Value() string {
	return c.value
}

func (c UpdateIpsecPolicyEncryptionAlgorithm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateIpsecPolicyEncryptionAlgorithm) UnmarshalJSON(b []byte) error {
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

type UpdateIpsecPolicyTransformProtocol struct {
	value string
}

type UpdateIpsecPolicyTransformProtocolEnum struct {
	ESP UpdateIpsecPolicyTransformProtocol
}

func GetUpdateIpsecPolicyTransformProtocolEnum() UpdateIpsecPolicyTransformProtocolEnum {
	return UpdateIpsecPolicyTransformProtocolEnum{
		ESP: UpdateIpsecPolicyTransformProtocol{
			value: "esp",
		},
	}
}

func (c UpdateIpsecPolicyTransformProtocol) Value() string {
	return c.value
}

func (c UpdateIpsecPolicyTransformProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateIpsecPolicyTransformProtocol) UnmarshalJSON(b []byte) error {
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

type UpdateIpsecPolicyEncapsulationMode struct {
	value string
}

type UpdateIpsecPolicyEncapsulationModeEnum struct {
	TUNNEL UpdateIpsecPolicyEncapsulationMode
}

func GetUpdateIpsecPolicyEncapsulationModeEnum() UpdateIpsecPolicyEncapsulationModeEnum {
	return UpdateIpsecPolicyEncapsulationModeEnum{
		TUNNEL: UpdateIpsecPolicyEncapsulationMode{
			value: "tunnel",
		},
	}
}

func (c UpdateIpsecPolicyEncapsulationMode) Value() string {
	return c.value
}

func (c UpdateIpsecPolicyEncapsulationMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateIpsecPolicyEncapsulationMode) UnmarshalJSON(b []byte) error {
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
