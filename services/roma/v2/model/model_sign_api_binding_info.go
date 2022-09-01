package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type SignApiBindingInfo struct {

	// API的发布编号
	PublishId *string `json:"publish_id,omitempty" xml:"publish_id"`

	// API编号
	ApiId *string `json:"api_id,omitempty" xml:"api_id"`

	// API所属分组的名称
	GroupName *string `json:"group_name,omitempty" xml:"group_name"`

	// 绑定时间
	BindingTime *sdktime.SdkTime `json:"binding_time,omitempty" xml:"binding_time"`

	// API所属环境的编号
	EnvId *string `json:"env_id,omitempty" xml:"env_id"`

	// API所属环境的名称
	EnvName *string `json:"env_name,omitempty" xml:"env_name"`

	// API类型
	ApiType *int32 `json:"api_type,omitempty" xml:"api_type"`

	// API名称
	ApiName *string `json:"api_name,omitempty" xml:"api_name"`

	// 绑定关系的ID
	Id *string `json:"id,omitempty" xml:"id"`

	// API描述
	ApiRemark *string `json:"api_remark,omitempty" xml:"api_remark"`

	// 签名密钥的编号
	SignId *string `json:"sign_id,omitempty" xml:"sign_id"`

	// 签名密钥的名称。支持汉字，英文，数字，下划线，且只能以英文和汉字开头，3 ~ 64字符。 > 中文字符必须为UTF-8或者unicode编码。
	SignName *string `json:"sign_name,omitempty" xml:"sign_name"`

	// 签名密钥的key。 - hmac类型的签名密钥key：支持英文，数字，下划线，中划线，且只能以英文字母或数字开头，8 ~ 32字符。未填写时后台自动生成。 - basic类型的签名密钥key：支持英文，数字，下划线，中划线，且只能以英文字母开头，4 ~ 32字符。未填写时后台自动生成。 - public_key类型的签名密钥key：支持英文，数字，下划线，中划线，+，/，=，可以英文字母，数字，+，/开头，8 ~ 512字符。未填写时后台自动生成。 - aes类型的签名秘钥key：支持英文，数字，下划线，中划线，!，@，#，$，%，+，/，=，可以英文字母，数字，+，/开头，签名算法为aes-128-cfb时为16个字符，签名算法为aes-256-cfb时为32个字符。未填写时后台自动生成。
	SignKey *string `json:"sign_key,omitempty" xml:"sign_key"`

	// 签名密钥的密钥。 - hmac类型的签名密钥key：支持英文，数字，下划线，中划线，!，@，#，$，%，且只能以英文字母或数字开头，16 ~ 64字符。未填写时后台自动生成。 - basic类型的签名密钥key：支持英文，数字，下划线，中划线，!，@，#，$，%，且只能以英文字母或数字开头，8 ~ 64字符。未填写时后台自动生成。 - public_key类型的签名密钥key：支持英文，数字，下划线，中划线，!，@，#，$，%，+，/，=，可以英文字母，数字，+，/开头，15 ~ 2048字符。未填写时后台自动生成。 - aes类型签名秘钥使用的向量：支持英文，数字，下划线，中划线，!，@，#，$，%，+，/，=，可以英文字母，数字，+，/开头，16个字符。未填写时后台自动生成。
	SignSecret *string `json:"sign_secret,omitempty" xml:"sign_secret"`

	// 签名密钥类型： - hmac - basic - public_key - aes  basic类型需要实例升级到对应版本，若不存在可联系技术工程师升级。  public_key类型开启实例配置public_key才可使用，实例特性配置详情请参考“附录 > 实例支持的APIC特性”，如确认实例不存在public_key配置可联系技术工程师开启。  aes类型需要实例升级到对应版本，若不存在可联系技术工程师升级。
	SignType *SignApiBindingInfoSignType `json:"sign_type,omitempty" xml:"sign_type"`
}

func (o SignApiBindingInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SignApiBindingInfo struct{}"
	}

	return strings.Join([]string{"SignApiBindingInfo", string(data)}, " ")
}

type SignApiBindingInfoSignType struct {
	value string
}

type SignApiBindingInfoSignTypeEnum struct {
	HMAC       SignApiBindingInfoSignType
	BASIC      SignApiBindingInfoSignType
	PUBLIC_KEY SignApiBindingInfoSignType
	AES        SignApiBindingInfoSignType
}

func GetSignApiBindingInfoSignTypeEnum() SignApiBindingInfoSignTypeEnum {
	return SignApiBindingInfoSignTypeEnum{
		HMAC: SignApiBindingInfoSignType{
			value: "hmac",
		},
		BASIC: SignApiBindingInfoSignType{
			value: "basic",
		},
		PUBLIC_KEY: SignApiBindingInfoSignType{
			value: "public_key",
		},
		AES: SignApiBindingInfoSignType{
			value: "aes",
		},
	}
}

func (c SignApiBindingInfoSignType) Value() string {
	return c.value
}

func (c SignApiBindingInfoSignType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SignApiBindingInfoSignType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
