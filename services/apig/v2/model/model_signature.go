package model

import (
	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type Signature struct {

	// 签名密钥的名称。支持汉字，英文，数字，下划线，且只能以英文和汉字开头。 > 中文字符必须为UTF-8或者unicode编码。
	Name string `json:"name"`

	// 签名密钥类型： - hmac - basic - public_key - aes  basic类型需要实例升级到对应版本，若不存在可联系技术工程师升级。  public_key类型开启实例配置public_key才可使用，实例特性配置详情请参考“附录 > 实例支持的APIG特性”，如确认实例不存在public_key配置可联系技术工程师开启。  aes类型需要实例升级到对应版本，若不存在可联系技术工程师升级。
	SignType *SignatureSignType `json:"sign_type,omitempty"`

	// 签名密钥的key。 - hmac类型的签名密钥key：支持英文，数字，下划线，中划线，且只能以英文字母或数字开头，8 ~ 32字符。未填写时后台自动生成。 - basic类型的签名密钥key：支持英文，数字，下划线，中划线，且只能以英文字母开头，4 ~ 32字符。未填写时后台自动生成。 - public_key类型的签名密钥key：支持英文，数字，下划线，中划线，+，/，=，可以英文字母，数字，+，/开头，8 ~ 512字符。未填写时后台自动生成。 - aes类型的签名密钥key：支持英文，数字，下划线，中划线，!，@，#，$，%，+，/，=，可以英文字母，数字，+，/开头，签名算法为aes-128-cfb时为16个字符，签名算法为aes-256-cfb时为32个字符。未填写时后台自动生成。
	SignKey *string `json:"sign_key,omitempty"`

	// 签名密钥的密钥。 - hmac类型的签名密钥key：支持英文，数字，下划线，中划线，!，@，#，$，%，且只能以英文字母或数字开头，16 ~ 64字符。未填写时后台自动生成。 - basic类型的签名密钥key：支持英文，数字，下划线，中划线，!，@，#，$，%，且只能以英文字母或数字开头，8 ~ 64字符。未填写时后台自动生成。 - public_key类型的签名密钥key：支持英文，数字，下划线，中划线，!，@，#，$，%，+，/，=，可以英文字母，数字，+，/开头，15 ~ 2048字符。未填写时后台自动生成。 - aes类型签名密钥使用的向量：支持英文，数字，下划线，中划线，!，@，#，$，%，+，/，=，可以英文字母，数字，+，/开头，16个字符。未填写时后台自动生成。
	SignSecret *string `json:"sign_secret,omitempty"`

	// 签名算法。默认值为空，仅aes类型签名密钥支持选择签名算法，其他类型签名密钥不支持签名算法。
	SignAlgorithm *SignatureSignAlgorithm `json:"sign_algorithm,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 签名密钥的编号
	Id *string `json:"id,omitempty"`
}

func (o Signature) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Signature struct{}"
	}

	return strings.Join([]string{"Signature", string(data)}, " ")
}

type SignatureSignType struct {
	value string
}

type SignatureSignTypeEnum struct {
	HMAC       SignatureSignType
	BASIC      SignatureSignType
	PUBLIC_KEY SignatureSignType
	AES        SignatureSignType
}

func GetSignatureSignTypeEnum() SignatureSignTypeEnum {
	return SignatureSignTypeEnum{
		HMAC: SignatureSignType{
			value: "hmac",
		},
		BASIC: SignatureSignType{
			value: "basic",
		},
		PUBLIC_KEY: SignatureSignType{
			value: "public_key",
		},
		AES: SignatureSignType{
			value: "aes",
		},
	}
}

func (c SignatureSignType) Value() string {
	return c.value
}

func (c SignatureSignType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SignatureSignType) UnmarshalJSON(b []byte) error {
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

type SignatureSignAlgorithm struct {
	value string
}

type SignatureSignAlgorithmEnum struct {
	AES_128_CFB SignatureSignAlgorithm
	AES_256_CFB SignatureSignAlgorithm
}

func GetSignatureSignAlgorithmEnum() SignatureSignAlgorithmEnum {
	return SignatureSignAlgorithmEnum{
		AES_128_CFB: SignatureSignAlgorithm{
			value: "aes-128-cfb",
		},
		AES_256_CFB: SignatureSignAlgorithm{
			value: "aes-256-cfb",
		},
	}
}

func (c SignatureSignAlgorithm) Value() string {
	return c.value
}

func (c SignatureSignAlgorithm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SignatureSignAlgorithm) UnmarshalJSON(b []byte) error {
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
