package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateServerRequestSslOptions SSL隧道协议的可选配置项。隧道协议类型是SSL时，必填
type CreateServerRequestSslOptions struct {

	// 协议
	Protocol *CreateServerRequestSslOptionsProtocol `json:"protocol,omitempty"`

	// 端口
	Port *CreateServerRequestSslOptionsPort `json:"port,omitempty"`

	// 加密算法
	EncryptionAlgorithm *CreateServerRequestSslOptionsEncryptionAlgorithm `json:"encryption_algorithm,omitempty"`

	// 是否压缩
	IsCompressed *bool `json:"is_compressed,omitempty"`
}

func (o CreateServerRequestSslOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServerRequestSslOptions struct{}"
	}

	return strings.Join([]string{"CreateServerRequestSslOptions", string(data)}, " ")
}

type CreateServerRequestSslOptionsProtocol struct {
	value string
}

type CreateServerRequestSslOptionsProtocolEnum struct {
	TCP CreateServerRequestSslOptionsProtocol
}

func GetCreateServerRequestSslOptionsProtocolEnum() CreateServerRequestSslOptionsProtocolEnum {
	return CreateServerRequestSslOptionsProtocolEnum{
		TCP: CreateServerRequestSslOptionsProtocol{
			value: "TCP",
		},
	}
}

func (c CreateServerRequestSslOptionsProtocol) Value() string {
	return c.value
}

func (c CreateServerRequestSslOptionsProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateServerRequestSslOptionsProtocol) UnmarshalJSON(b []byte) error {
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

type CreateServerRequestSslOptionsPort struct {
	value int32
}

type CreateServerRequestSslOptionsPortEnum struct {
	E_443  CreateServerRequestSslOptionsPort
	E_1194 CreateServerRequestSslOptionsPort
}

func GetCreateServerRequestSslOptionsPortEnum() CreateServerRequestSslOptionsPortEnum {
	return CreateServerRequestSslOptionsPortEnum{
		E_443: CreateServerRequestSslOptionsPort{
			value: 443,
		}, E_1194: CreateServerRequestSslOptionsPort{
			value: 1194,
		},
	}
}

func (c CreateServerRequestSslOptionsPort) Value() int32 {
	return c.value
}

func (c CreateServerRequestSslOptionsPort) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateServerRequestSslOptionsPort) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type CreateServerRequestSslOptionsEncryptionAlgorithm struct {
	value string
}

type CreateServerRequestSslOptionsEncryptionAlgorithmEnum struct {
	AES_128_GCM CreateServerRequestSslOptionsEncryptionAlgorithm
	AES_256_GCM CreateServerRequestSslOptionsEncryptionAlgorithm
}

func GetCreateServerRequestSslOptionsEncryptionAlgorithmEnum() CreateServerRequestSslOptionsEncryptionAlgorithmEnum {
	return CreateServerRequestSslOptionsEncryptionAlgorithmEnum{
		AES_128_GCM: CreateServerRequestSslOptionsEncryptionAlgorithm{
			value: "AES-128-GCM",
		},
		AES_256_GCM: CreateServerRequestSslOptionsEncryptionAlgorithm{
			value: "AES-256-GCM",
		},
	}
}

func (c CreateServerRequestSslOptionsEncryptionAlgorithm) Value() string {
	return c.value
}

func (c CreateServerRequestSslOptionsEncryptionAlgorithm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateServerRequestSslOptionsEncryptionAlgorithm) UnmarshalJSON(b []byte) error {
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
