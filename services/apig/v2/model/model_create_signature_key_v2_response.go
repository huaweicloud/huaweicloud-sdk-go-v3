package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// CreateSignatureKeyV2Response Response Object
type CreateSignatureKeyV2Response struct {

	// 签名密钥的名称。支持汉字，英文，数字，下划线，且只能以英文和汉字开头。 > 中文字符必须为UTF-8或者unicode编码。
	Name string `json:"name"`

	// 签名密钥类型： - hmac - basic - public_key - aes  basic类型需要实例升级到对应版本，如果不存在可联系技术工程师升级。  public_key类型开启实例配置public_key才可使用，实例特性配置详情请参考“附录 > 实例支持的APIG特性”，如确认实例不存在public_key配置可联系技术工程师开启。  aes类型需要实例升级到对应版本，如果不存在可联系技术工程师升级。
	SignType *CreateSignatureKeyV2ResponseSignType `json:"sign_type,omitempty"`

	// 签名密钥的key。 - hmac类型的签名密钥key：支持英文，数字，下划线，中划线，且只能以英文字母或数字开头，8 ~ 32字符。未填写时后台自动生成。 - basic类型的签名密钥key：支持英文，数字，下划线，中划线，且只能以英文字母开头，4 ~ 32字符。未填写时后台自动生成。 - public_key类型的签名密钥key：支持英文，数字，下划线，中划线，+，/，=，可以英文字母，数字，+，/开头，8 ~ 512字符。未填写时后台自动生成。 - aes类型的签名密钥key：支持英文，数字，下划线，中划线，!，@，#，$，%，+，/，=，可以英文字母，数字，+，/开头，签名算法为aes-128-cfb时为16个字符，签名算法为aes-256-cfb时为32个字符。未填写时后台自动生成。
	SignKey *string `json:"sign_key,omitempty"`

	// 签名密钥的密钥。 - hmac类型的签名密钥key：支持英文，数字，下划线，中划线，!，@，#，$，%，且只能以英文字母或数字开头，16 ~ 64字符。未填写时后台自动生成。 - basic类型的签名密钥key：支持英文，数字，下划线，中划线，!，@，#，$，%，且只能以英文字母或数字开头，8 ~ 64字符。未填写时后台自动生成。 - public_key类型的签名密钥key：支持英文，数字，下划线，中划线，!，@，#，$，%，+，/，=，可以英文字母，数字，+，/开头，16 ~ 2048字符。未填写时后台自动生成。 - aes类型签名密钥使用的向量：支持英文，数字，下划线，中划线，!，@，#，$，%，+，/，=，可以英文字母，数字，+，/开头，16个字符。未填写时后台自动生成。
	SignSecret *string `json:"sign_secret,omitempty"`

	// 签名算法。默认值为空，仅aes类型签名密钥支持选择签名算法，其他类型签名密钥不支持签名算法。
	SignAlgorithm *CreateSignatureKeyV2ResponseSignAlgorithm `json:"sign_algorithm,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 签名密钥的编号
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSignatureKeyV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSignatureKeyV2Response struct{}"
	}

	return strings.Join([]string{"CreateSignatureKeyV2Response", string(data)}, " ")
}

type CreateSignatureKeyV2ResponseSignType struct {
	value string
}

type CreateSignatureKeyV2ResponseSignTypeEnum struct {
	HMAC       CreateSignatureKeyV2ResponseSignType
	BASIC      CreateSignatureKeyV2ResponseSignType
	PUBLIC_KEY CreateSignatureKeyV2ResponseSignType
	AES        CreateSignatureKeyV2ResponseSignType
}

func GetCreateSignatureKeyV2ResponseSignTypeEnum() CreateSignatureKeyV2ResponseSignTypeEnum {
	return CreateSignatureKeyV2ResponseSignTypeEnum{
		HMAC: CreateSignatureKeyV2ResponseSignType{
			value: "hmac",
		},
		BASIC: CreateSignatureKeyV2ResponseSignType{
			value: "basic",
		},
		PUBLIC_KEY: CreateSignatureKeyV2ResponseSignType{
			value: "public_key",
		},
		AES: CreateSignatureKeyV2ResponseSignType{
			value: "aes",
		},
	}
}

func (c CreateSignatureKeyV2ResponseSignType) Value() string {
	return c.value
}

func (c CreateSignatureKeyV2ResponseSignType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSignatureKeyV2ResponseSignType) UnmarshalJSON(b []byte) error {
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

type CreateSignatureKeyV2ResponseSignAlgorithm struct {
	value string
}

type CreateSignatureKeyV2ResponseSignAlgorithmEnum struct {
	AES_128_CFB CreateSignatureKeyV2ResponseSignAlgorithm
	AES_256_CFB CreateSignatureKeyV2ResponseSignAlgorithm
}

func GetCreateSignatureKeyV2ResponseSignAlgorithmEnum() CreateSignatureKeyV2ResponseSignAlgorithmEnum {
	return CreateSignatureKeyV2ResponseSignAlgorithmEnum{
		AES_128_CFB: CreateSignatureKeyV2ResponseSignAlgorithm{
			value: "aes-128-cfb",
		},
		AES_256_CFB: CreateSignatureKeyV2ResponseSignAlgorithm{
			value: "aes-256-cfb",
		},
	}
}

func (c CreateSignatureKeyV2ResponseSignAlgorithm) Value() string {
	return c.value
}

func (c CreateSignatureKeyV2ResponseSignAlgorithm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSignatureKeyV2ResponseSignAlgorithm) UnmarshalJSON(b []byte) error {
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
