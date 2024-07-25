package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type SignatureWithBindNum struct {

	// 签名密钥的名称。支持汉字，英文，数字，下划线，且只能以英文和汉字开头。 > 中文字符必须为UTF-8或者unicode编码。
	Name string `json:"name"`

	// 签名密钥类型： - hmac - basic - public_key - aes  basic类型需要实例升级到对应版本，如果不存在可联系技术工程师升级。  public_key类型开启实例配置public_key才可使用，实例特性配置详情请参考“附录 > 实例支持的APIG特性”，如确认实例不存在public_key配置可联系技术工程师开启。  aes类型需要实例升级到对应版本，如果不存在可联系技术工程师升级。
	SignType *SignatureWithBindNumSignType `json:"sign_type,omitempty"`

	// 签名密钥的key。 - hmac类型的签名密钥key：支持英文，数字，下划线，中划线，且只能以英文字母或数字开头，8 ~ 32字符。未填写时后台自动生成。 - basic类型的签名密钥key：支持英文，数字，下划线，中划线，且只能以英文字母开头，4 ~ 32字符。未填写时后台自动生成。 - public_key类型的签名密钥key：支持英文，数字，下划线，中划线，+，/，=，可以英文字母，数字，+，/开头，8 ~ 512字符。未填写时后台自动生成。 - aes类型的签名密钥key：支持英文，数字，下划线，中划线，!，@，#，$，%，+，/，=，可以英文字母，数字，+，/开头，签名算法为aes-128-cfb时为16个字符，签名算法为aes-256-cfb时为32个字符。未填写时后台自动生成。
	SignKey *string `json:"sign_key,omitempty"`

	// 签名密钥的密钥。 - hmac类型的签名密钥key：支持英文，数字，下划线，中划线，!，@，#，$，%，且只能以英文字母或数字开头，16 ~ 64字符。未填写时后台自动生成。 - basic类型的签名密钥key：支持英文，数字，下划线，中划线，!，@，#，$，%，且只能以英文字母或数字开头，8 ~ 64字符。未填写时后台自动生成。 - public_key类型的签名密钥key：支持英文，数字，下划线，中划线，!，@，#，$，%，+，/，=，可以英文字母，数字，+，/开头，16 ~ 2048字符。未填写时后台自动生成。 - aes类型签名密钥使用的向量：支持英文，数字，下划线，中划线，!，@，#，$，%，+，/，=，可以英文字母，数字，+，/开头，16个字符。未填写时后台自动生成。
	SignSecret *string `json:"sign_secret,omitempty"`

	// 签名算法。默认值为空，仅aes类型签名密钥支持选择签名算法，其他类型签名密钥不支持签名算法。
	SignAlgorithm *SignatureWithBindNumSignAlgorithm `json:"sign_algorithm,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 签名密钥的编号
	Id *string `json:"id,omitempty"`

	// 绑定的API数量
	BindNum *int32 `json:"bind_num,omitempty"`

	// 绑定的自定义后端数量  暂不支持
	LdapiBindNum *int32 `json:"ldapi_bind_num,omitempty"`
}

func (o SignatureWithBindNum) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SignatureWithBindNum struct{}"
	}

	return strings.Join([]string{"SignatureWithBindNum", string(data)}, " ")
}

type SignatureWithBindNumSignType struct {
	value string
}

type SignatureWithBindNumSignTypeEnum struct {
	HMAC       SignatureWithBindNumSignType
	BASIC      SignatureWithBindNumSignType
	PUBLIC_KEY SignatureWithBindNumSignType
	AES        SignatureWithBindNumSignType
}

func GetSignatureWithBindNumSignTypeEnum() SignatureWithBindNumSignTypeEnum {
	return SignatureWithBindNumSignTypeEnum{
		HMAC: SignatureWithBindNumSignType{
			value: "hmac",
		},
		BASIC: SignatureWithBindNumSignType{
			value: "basic",
		},
		PUBLIC_KEY: SignatureWithBindNumSignType{
			value: "public_key",
		},
		AES: SignatureWithBindNumSignType{
			value: "aes",
		},
	}
}

func (c SignatureWithBindNumSignType) Value() string {
	return c.value
}

func (c SignatureWithBindNumSignType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SignatureWithBindNumSignType) UnmarshalJSON(b []byte) error {
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

type SignatureWithBindNumSignAlgorithm struct {
	value string
}

type SignatureWithBindNumSignAlgorithmEnum struct {
	AES_128_CFB SignatureWithBindNumSignAlgorithm
	AES_256_CFB SignatureWithBindNumSignAlgorithm
}

func GetSignatureWithBindNumSignAlgorithmEnum() SignatureWithBindNumSignAlgorithmEnum {
	return SignatureWithBindNumSignAlgorithmEnum{
		AES_128_CFB: SignatureWithBindNumSignAlgorithm{
			value: "aes-128-cfb",
		},
		AES_256_CFB: SignatureWithBindNumSignAlgorithm{
			value: "aes-256-cfb",
		},
	}
}

func (c SignatureWithBindNumSignAlgorithm) Value() string {
	return c.value
}

func (c SignatureWithBindNumSignAlgorithm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SignatureWithBindNumSignAlgorithm) UnmarshalJSON(b []byte) error {
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
