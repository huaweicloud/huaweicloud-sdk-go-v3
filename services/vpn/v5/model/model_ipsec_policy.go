package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type IpsecPolicy struct {

	// 认证算法，SHA1和MD5安全性较低，请慎用
	AuthenticationAlgorithm *IpsecPolicyAuthenticationAlgorithm `json:"authentication_algorithm,omitempty"`

	// 加密算法，3DES安全性较低，请慎用
	EncryptionAlgorithm *IpsecPolicyEncryptionAlgorithm `json:"encryption_algorithm,omitempty"`

	// PFS使用的DH密钥组
	Pfs *string `json:"pfs,omitempty"`

	// 传输模式
	TransformProtocol *IpsecPolicyTransformProtocol `json:"transform_protocol,omitempty"`

	// 表示配置IPSec连接建立的隧道以时间为基准的生存周期
	LifetimeSeconds *int32 `json:"lifetime_seconds,omitempty"`

	// 封装模式，当前只有tunnel模式
	EncapsulationMode *IpsecPolicyEncapsulationMode `json:"encapsulation_mode,omitempty"`
}

func (o IpsecPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpsecPolicy struct{}"
	}

	return strings.Join([]string{"IpsecPolicy", string(data)}, " ")
}

type IpsecPolicyAuthenticationAlgorithm struct {
	value string
}

type IpsecPolicyAuthenticationAlgorithmEnum struct {
	SHA1     IpsecPolicyAuthenticationAlgorithm
	MD5      IpsecPolicyAuthenticationAlgorithm
	SHA2_256 IpsecPolicyAuthenticationAlgorithm
	SHA2_384 IpsecPolicyAuthenticationAlgorithm
	SHA2_512 IpsecPolicyAuthenticationAlgorithm
}

func GetIpsecPolicyAuthenticationAlgorithmEnum() IpsecPolicyAuthenticationAlgorithmEnum {
	return IpsecPolicyAuthenticationAlgorithmEnum{
		SHA1: IpsecPolicyAuthenticationAlgorithm{
			value: "sha1",
		},
		MD5: IpsecPolicyAuthenticationAlgorithm{
			value: "md5",
		},
		SHA2_256: IpsecPolicyAuthenticationAlgorithm{
			value: "sha2-256",
		},
		SHA2_384: IpsecPolicyAuthenticationAlgorithm{
			value: "sha2-384",
		},
		SHA2_512: IpsecPolicyAuthenticationAlgorithm{
			value: "sha2-512",
		},
	}
}

func (c IpsecPolicyAuthenticationAlgorithm) Value() string {
	return c.value
}

func (c IpsecPolicyAuthenticationAlgorithm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IpsecPolicyAuthenticationAlgorithm) UnmarshalJSON(b []byte) error {
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

type IpsecPolicyEncryptionAlgorithm struct {
	value string
}

type IpsecPolicyEncryptionAlgorithmEnum struct {
	E_3DES          IpsecPolicyEncryptionAlgorithm
	AES_128         IpsecPolicyEncryptionAlgorithm
	AES_192         IpsecPolicyEncryptionAlgorithm
	AES_256         IpsecPolicyEncryptionAlgorithm
	AES_128_GCM_16  IpsecPolicyEncryptionAlgorithm
	AES_256_GCM_16  IpsecPolicyEncryptionAlgorithm
	AES_128_GCM_128 IpsecPolicyEncryptionAlgorithm
	AES_256_GCM_128 IpsecPolicyEncryptionAlgorithm
}

func GetIpsecPolicyEncryptionAlgorithmEnum() IpsecPolicyEncryptionAlgorithmEnum {
	return IpsecPolicyEncryptionAlgorithmEnum{
		E_3DES: IpsecPolicyEncryptionAlgorithm{
			value: "3des",
		},
		AES_128: IpsecPolicyEncryptionAlgorithm{
			value: "aes-128",
		},
		AES_192: IpsecPolicyEncryptionAlgorithm{
			value: "aes-192",
		},
		AES_256: IpsecPolicyEncryptionAlgorithm{
			value: "aes-256",
		},
		AES_128_GCM_16: IpsecPolicyEncryptionAlgorithm{
			value: "aes-128-gcm-16",
		},
		AES_256_GCM_16: IpsecPolicyEncryptionAlgorithm{
			value: "aes-256-gcm-16",
		},
		AES_128_GCM_128: IpsecPolicyEncryptionAlgorithm{
			value: "aes-128-gcm-128",
		},
		AES_256_GCM_128: IpsecPolicyEncryptionAlgorithm{
			value: "aes-256-gcm-128",
		},
	}
}

func (c IpsecPolicyEncryptionAlgorithm) Value() string {
	return c.value
}

func (c IpsecPolicyEncryptionAlgorithm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IpsecPolicyEncryptionAlgorithm) UnmarshalJSON(b []byte) error {
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

type IpsecPolicyTransformProtocol struct {
	value string
}

type IpsecPolicyTransformProtocolEnum struct {
	ESP IpsecPolicyTransformProtocol
}

func GetIpsecPolicyTransformProtocolEnum() IpsecPolicyTransformProtocolEnum {
	return IpsecPolicyTransformProtocolEnum{
		ESP: IpsecPolicyTransformProtocol{
			value: "esp",
		},
	}
}

func (c IpsecPolicyTransformProtocol) Value() string {
	return c.value
}

func (c IpsecPolicyTransformProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IpsecPolicyTransformProtocol) UnmarshalJSON(b []byte) error {
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

type IpsecPolicyEncapsulationMode struct {
	value string
}

type IpsecPolicyEncapsulationModeEnum struct {
	TUNNEL IpsecPolicyEncapsulationMode
}

func GetIpsecPolicyEncapsulationModeEnum() IpsecPolicyEncapsulationModeEnum {
	return IpsecPolicyEncapsulationModeEnum{
		TUNNEL: IpsecPolicyEncapsulationMode{
			value: "tunnel",
		},
	}
}

func (c IpsecPolicyEncapsulationMode) Value() string {
	return c.value
}

func (c IpsecPolicyEncapsulationMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IpsecPolicyEncapsulationMode) UnmarshalJSON(b []byte) error {
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
