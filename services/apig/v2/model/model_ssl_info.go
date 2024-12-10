package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SslInfo struct {

	// SSL证书编号。
	SslId *string `json:"ssl_id,omitempty"`

	// SSL证书名称。
	SslName *string `json:"ssl_name,omitempty"`

	// 证书算法类型： - RSA - ECC - SM2
	AlgorithmType *SslInfoAlgorithmType `json:"algorithm_type,omitempty"`

	// 证书可见范围： - instance：当前实例 - global：全局
	Type *SslInfoType `json:"type,omitempty"`
}

func (o SslInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SslInfo struct{}"
	}

	return strings.Join([]string{"SslInfo", string(data)}, " ")
}

type SslInfoAlgorithmType struct {
	value string
}

type SslInfoAlgorithmTypeEnum struct {
	RSA SslInfoAlgorithmType
	ECC SslInfoAlgorithmType
	SM2 SslInfoAlgorithmType
}

func GetSslInfoAlgorithmTypeEnum() SslInfoAlgorithmTypeEnum {
	return SslInfoAlgorithmTypeEnum{
		RSA: SslInfoAlgorithmType{
			value: "RSA",
		},
		ECC: SslInfoAlgorithmType{
			value: "ECC",
		},
		SM2: SslInfoAlgorithmType{
			value: "SM2",
		},
	}
}

func (c SslInfoAlgorithmType) Value() string {
	return c.value
}

func (c SslInfoAlgorithmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SslInfoAlgorithmType) UnmarshalJSON(b []byte) error {
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

type SslInfoType struct {
	value string
}

type SslInfoTypeEnum struct {
	INSTANCE SslInfoType
	GLOBAL   SslInfoType
}

func GetSslInfoTypeEnum() SslInfoTypeEnum {
	return SslInfoTypeEnum{
		INSTANCE: SslInfoType{
			value: "instance",
		},
		GLOBAL: SslInfoType{
			value: "global",
		},
	}
}

func (c SslInfoType) Value() string {
	return c.value
}

func (c SslInfoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SslInfoType) UnmarshalJSON(b []byte) error {
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
