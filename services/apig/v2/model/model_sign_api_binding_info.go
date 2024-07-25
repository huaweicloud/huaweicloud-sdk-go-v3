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
	PublishId *string `json:"publish_id,omitempty"`

	// API编号
	ApiId *string `json:"api_id,omitempty"`

	// API所属分组的名称
	GroupName *string `json:"group_name,omitempty"`

	// 绑定时间
	BindingTime *sdktime.SdkTime `json:"binding_time,omitempty"`

	// API所属环境的编号
	EnvId *string `json:"env_id,omitempty"`

	// API所属环境的名称
	EnvName *string `json:"env_name,omitempty"`

	// API类型
	ApiType *int32 `json:"api_type,omitempty"`

	// API名称
	ApiName *string `json:"api_name,omitempty"`

	// 绑定关系的ID
	Id *string `json:"id,omitempty"`

	// API描述
	ApiRemark *string `json:"api_remark,omitempty"`

	// 签名密钥的编号
	SignId *string `json:"sign_id,omitempty"`

	// 签名密钥的名称。支持汉字，英文，数字，下划线，且只能以英文和汉字开头，3 ~ 64字符。 > 中文字符必须为UTF-8或者unicode编码。
	SignName *string `json:"sign_name,omitempty"`

	// API请求方法
	ReqMethod *SignApiBindingInfoReqMethod `json:"req_method,omitempty"`

	// 签名密钥的key。 - hmac类型的签名密钥key：支持英文，数字，下划线，中划线，且只能以英文字母或数字开头，8 ~ 32字符。未填写时后台自动生成。 - basic类型的签名密钥key：支持英文，数字，下划线，中划线，且只能以英文字母开头，4 ~ 32字符。未填写时后台自动生成。 - public_key类型的签名密钥key：支持英文，数字，下划线，中划线，+，/，=，可以英文字母，数字，+，/开头，8 ~ 512字符。未填写时后台自动生成。 - aes类型的签名密钥key：支持英文，数字，下划线，中划线，!，@，#，$，%，+，/，=，可以英文字母，数字，+，/开头，签名算法为aes-128-cfb时为16个字符，签名算法为aes-256-cfb时为32个字符。未填写时后台自动生成。
	SignKey *string `json:"sign_key,omitempty"`

	// 签名密钥的密钥。 - hmac类型的签名密钥key：支持英文，数字，下划线，中划线，!，@，#，$，%，且只能以英文字母或数字开头，16 ~ 64字符。未填写时后台自动生成。 - basic类型的签名密钥key：支持英文，数字，下划线，中划线，!，@，#，$，%，且只能以英文字母或数字开头，8 ~ 64字符。未填写时后台自动生成。 - public_key类型的签名密钥key：支持英文，数字，下划线，中划线，!，@，#，$，%，+，/，=，可以英文字母，数字，+，/开头，15 ~ 2048字符。未填写时后台自动生成。 - aes类型签名密钥使用的向量：支持英文，数字，下划线，中划线，!，@，#，$，%，+，/，=，可以英文字母，数字，+，/开头，16个字符。未填写时后台自动生成。
	SignSecret *string `json:"sign_secret,omitempty"`

	// 签名密钥类型： - hmac - basic - public_key - aes  basic类型需要实例升级到对应版本，如果不存在可联系技术工程师升级。  public_key类型开启实例配置public_key才可使用，实例特性配置详情请参考“附录 > 实例支持的APIG特性”，如确认实例不存在public_key配置可联系技术工程师开启。  aes类型需要实例升级到对应版本，如果不存在可联系技术工程师升级。
	SignType *SignApiBindingInfoSignType `json:"sign_type,omitempty"`
}

func (o SignApiBindingInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SignApiBindingInfo struct{}"
	}

	return strings.Join([]string{"SignApiBindingInfo", string(data)}, " ")
}

type SignApiBindingInfoReqMethod struct {
	value string
}

type SignApiBindingInfoReqMethodEnum struct {
	GET     SignApiBindingInfoReqMethod
	POST    SignApiBindingInfoReqMethod
	DELETE  SignApiBindingInfoReqMethod
	PUT     SignApiBindingInfoReqMethod
	PATCH   SignApiBindingInfoReqMethod
	HEAD    SignApiBindingInfoReqMethod
	OPTIONS SignApiBindingInfoReqMethod
	ANY     SignApiBindingInfoReqMethod
}

func GetSignApiBindingInfoReqMethodEnum() SignApiBindingInfoReqMethodEnum {
	return SignApiBindingInfoReqMethodEnum{
		GET: SignApiBindingInfoReqMethod{
			value: "GET",
		},
		POST: SignApiBindingInfoReqMethod{
			value: "POST",
		},
		DELETE: SignApiBindingInfoReqMethod{
			value: "DELETE",
		},
		PUT: SignApiBindingInfoReqMethod{
			value: "PUT",
		},
		PATCH: SignApiBindingInfoReqMethod{
			value: "PATCH",
		},
		HEAD: SignApiBindingInfoReqMethod{
			value: "HEAD",
		},
		OPTIONS: SignApiBindingInfoReqMethod{
			value: "OPTIONS",
		},
		ANY: SignApiBindingInfoReqMethod{
			value: "ANY",
		},
	}
}

func (c SignApiBindingInfoReqMethod) Value() string {
	return c.value
}

func (c SignApiBindingInfoReqMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SignApiBindingInfoReqMethod) UnmarshalJSON(b []byte) error {
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
