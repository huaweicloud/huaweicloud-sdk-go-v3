package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FlowSourceDecryption StreamConnect flow解密信息
type FlowSourceDecryption struct {

	// 加密算法，aes128: 加密算法为aes-128，aes192:加密算法为aes-192，aes256: 加密算法为AES-256
	Algorithm *FlowSourceDecryptionAlgorithm `json:"algorithm,omitempty"`

	// 秘钥类型,speke:使用speke协议获取秘钥,static-key:静态秘钥,srt-password:SRT协议秘钥 目前仅支持srt-password类型，其他类型暂不支持
	KeyType FlowSourceDecryptionKeyType `json:"key_type"`

	// srt解密秘钥，用于flow对srt流进行解密
	Passphrase string `json:"passphrase"`
}

func (o FlowSourceDecryption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlowSourceDecryption struct{}"
	}

	return strings.Join([]string{"FlowSourceDecryption", string(data)}, " ")
}

type FlowSourceDecryptionAlgorithm struct {
	value string
}

type FlowSourceDecryptionAlgorithmEnum struct {
	AES128 FlowSourceDecryptionAlgorithm
	AES192 FlowSourceDecryptionAlgorithm
	AES256 FlowSourceDecryptionAlgorithm
}

func GetFlowSourceDecryptionAlgorithmEnum() FlowSourceDecryptionAlgorithmEnum {
	return FlowSourceDecryptionAlgorithmEnum{
		AES128: FlowSourceDecryptionAlgorithm{
			value: "aes128",
		},
		AES192: FlowSourceDecryptionAlgorithm{
			value: "aes192",
		},
		AES256: FlowSourceDecryptionAlgorithm{
			value: "aes256",
		},
	}
}

func (c FlowSourceDecryptionAlgorithm) Value() string {
	return c.value
}

func (c FlowSourceDecryptionAlgorithm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlowSourceDecryptionAlgorithm) UnmarshalJSON(b []byte) error {
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

type FlowSourceDecryptionKeyType struct {
	value string
}

type FlowSourceDecryptionKeyTypeEnum struct {
	SPEKE        FlowSourceDecryptionKeyType
	STATIC_KEY   FlowSourceDecryptionKeyType
	SRT_PASSWORD FlowSourceDecryptionKeyType
}

func GetFlowSourceDecryptionKeyTypeEnum() FlowSourceDecryptionKeyTypeEnum {
	return FlowSourceDecryptionKeyTypeEnum{
		SPEKE: FlowSourceDecryptionKeyType{
			value: "speke",
		},
		STATIC_KEY: FlowSourceDecryptionKeyType{
			value: "static-key",
		},
		SRT_PASSWORD: FlowSourceDecryptionKeyType{
			value: "srt-password",
		},
	}
}

func (c FlowSourceDecryptionKeyType) Value() string {
	return c.value
}

func (c FlowSourceDecryptionKeyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlowSourceDecryptionKeyType) UnmarshalJSON(b []byte) error {
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
