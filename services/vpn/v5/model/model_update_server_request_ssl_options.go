package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateServerRequestSslOptions SSL隧道协议的可选配置项
type UpdateServerRequestSslOptions struct {

	// 协议
	Protocol *UpdateServerRequestSslOptionsProtocol `json:"protocol,omitempty"`

	// 端口
	Port *UpdateServerRequestSslOptionsPort `json:"port,omitempty"`

	// 加密算法
	EncryptionAlgorithm *UpdateServerRequestSslOptionsEncryptionAlgorithm `json:"encryption_algorithm,omitempty"`
}

func (o UpdateServerRequestSslOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerRequestSslOptions struct{}"
	}

	return strings.Join([]string{"UpdateServerRequestSslOptions", string(data)}, " ")
}

type UpdateServerRequestSslOptionsProtocol struct {
	value string
}

type UpdateServerRequestSslOptionsProtocolEnum struct {
	TCP UpdateServerRequestSslOptionsProtocol
}

func GetUpdateServerRequestSslOptionsProtocolEnum() UpdateServerRequestSslOptionsProtocolEnum {
	return UpdateServerRequestSslOptionsProtocolEnum{
		TCP: UpdateServerRequestSslOptionsProtocol{
			value: "TCP",
		},
	}
}

func (c UpdateServerRequestSslOptionsProtocol) Value() string {
	return c.value
}

func (c UpdateServerRequestSslOptionsProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateServerRequestSslOptionsProtocol) UnmarshalJSON(b []byte) error {
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

type UpdateServerRequestSslOptionsPort struct {
	value int32
}

type UpdateServerRequestSslOptionsPortEnum struct {
	E_443  UpdateServerRequestSslOptionsPort
	E_1194 UpdateServerRequestSslOptionsPort
}

func GetUpdateServerRequestSslOptionsPortEnum() UpdateServerRequestSslOptionsPortEnum {
	return UpdateServerRequestSslOptionsPortEnum{
		E_443: UpdateServerRequestSslOptionsPort{
			value: 443,
		}, E_1194: UpdateServerRequestSslOptionsPort{
			value: 1194,
		},
	}
}

func (c UpdateServerRequestSslOptionsPort) Value() int32 {
	return c.value
}

func (c UpdateServerRequestSslOptionsPort) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateServerRequestSslOptionsPort) UnmarshalJSON(b []byte) error {
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

type UpdateServerRequestSslOptionsEncryptionAlgorithm struct {
	value string
}

type UpdateServerRequestSslOptionsEncryptionAlgorithmEnum struct {
	AES_128_GCM UpdateServerRequestSslOptionsEncryptionAlgorithm
	AES_256_GCM UpdateServerRequestSslOptionsEncryptionAlgorithm
}

func GetUpdateServerRequestSslOptionsEncryptionAlgorithmEnum() UpdateServerRequestSslOptionsEncryptionAlgorithmEnum {
	return UpdateServerRequestSslOptionsEncryptionAlgorithmEnum{
		AES_128_GCM: UpdateServerRequestSslOptionsEncryptionAlgorithm{
			value: "AES-128-GCM",
		},
		AES_256_GCM: UpdateServerRequestSslOptionsEncryptionAlgorithm{
			value: "AES-256-GCM",
		},
	}
}

func (c UpdateServerRequestSslOptionsEncryptionAlgorithm) Value() string {
	return c.value
}

func (c UpdateServerRequestSslOptionsEncryptionAlgorithm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateServerRequestSslOptionsEncryptionAlgorithm) UnmarshalJSON(b []byte) error {
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
