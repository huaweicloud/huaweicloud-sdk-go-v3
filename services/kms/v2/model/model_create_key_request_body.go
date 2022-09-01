package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateKeyRequestBody struct {

	// 非默认主密钥别名，取值范围为1到255个字符，满足正则匹配“^[a-zA-Z0-9:/_-]{1,255}$”，且不与系统服务创建的默认主密钥别名重名。
	KeyAlias string `json:"key_alias" xml:"key_alias"`

	// 密钥生成算法，默认为“AES_256”，枚举如下： - AES_256 - SM4 - RSA_2048 - RSA_3072 - RSA_4096 - EC_P256 - EC_P384 - SM2
	KeySpec *CreateKeyRequestBodyKeySpec `json:"key_spec,omitempty" xml:"key_spec"`

	// 密钥用途，对称密钥默认为“ENCRYPT_DECRYPT”，非对称密钥默认为“SIGN_VERIFY”，枚举如下： - ENCRYPT_DECRYPT - SIGN_VERIFY
	KeyUsage *CreateKeyRequestBodyKeyUsage `json:"key_usage,omitempty" xml:"key_usage"`

	// 密钥描述，取值0到255字符。
	KeyDescription *string `json:"key_description,omitempty" xml:"key_description"`

	// 密钥来源，默认为“kms”，枚举如下： - kms：表示密钥材料由kms生成。 - external：表示密钥材料由外部导入。
	Origin *CreateKeyRequestBodyOrigin `json:"origin,omitempty" xml:"origin"`

	// 企业多项目ID。 - 用户未开通企业多项目时，不需要输入该字段。 - 用户开通企业多项目时，创建资源可以输入该字段。若用户户不输入该字段，默认创建属于默认企业多项目ID（ID为“0”）的资源。 注意：若用户没有默认企业多项目ID（ID为“0”）下的创建权限，则接口报错。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	// 请求消息序列号，36字节序列号。 例如：919c82d4-8046-4722-9094-35c3c6524cff
	Sequence *string `json:"sequence,omitempty" xml:"sequence"`

	// 密钥库ID，默认使用KMS默认密钥库
	KeystoreId *string `json:"keystore_id,omitempty" xml:"keystore_id"`
}

func (o CreateKeyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKeyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateKeyRequestBody", string(data)}, " ")
}

type CreateKeyRequestBodyKeySpec struct {
	value string
}

type CreateKeyRequestBodyKeySpecEnum struct {
	AES_256  CreateKeyRequestBodyKeySpec
	SM4      CreateKeyRequestBodyKeySpec
	RSA_2048 CreateKeyRequestBodyKeySpec
	RSA_3072 CreateKeyRequestBodyKeySpec
	RSA_4096 CreateKeyRequestBodyKeySpec
	EC_P256  CreateKeyRequestBodyKeySpec
	EC_P384  CreateKeyRequestBodyKeySpec
	SM2      CreateKeyRequestBodyKeySpec
}

func GetCreateKeyRequestBodyKeySpecEnum() CreateKeyRequestBodyKeySpecEnum {
	return CreateKeyRequestBodyKeySpecEnum{
		AES_256: CreateKeyRequestBodyKeySpec{
			value: "AES_256",
		},
		SM4: CreateKeyRequestBodyKeySpec{
			value: "SM4",
		},
		RSA_2048: CreateKeyRequestBodyKeySpec{
			value: "RSA_2048",
		},
		RSA_3072: CreateKeyRequestBodyKeySpec{
			value: "RSA_3072",
		},
		RSA_4096: CreateKeyRequestBodyKeySpec{
			value: "RSA_4096",
		},
		EC_P256: CreateKeyRequestBodyKeySpec{
			value: "EC_P256",
		},
		EC_P384: CreateKeyRequestBodyKeySpec{
			value: "EC_P384",
		},
		SM2: CreateKeyRequestBodyKeySpec{
			value: "SM2",
		},
	}
}

func (c CreateKeyRequestBodyKeySpec) Value() string {
	return c.value
}

func (c CreateKeyRequestBodyKeySpec) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateKeyRequestBodyKeySpec) UnmarshalJSON(b []byte) error {
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

type CreateKeyRequestBodyKeyUsage struct {
	value string
}

type CreateKeyRequestBodyKeyUsageEnum struct {
	ENCRYPT_DECRYPT CreateKeyRequestBodyKeyUsage
	SIGN_VERIFY     CreateKeyRequestBodyKeyUsage
}

func GetCreateKeyRequestBodyKeyUsageEnum() CreateKeyRequestBodyKeyUsageEnum {
	return CreateKeyRequestBodyKeyUsageEnum{
		ENCRYPT_DECRYPT: CreateKeyRequestBodyKeyUsage{
			value: "ENCRYPT_DECRYPT",
		},
		SIGN_VERIFY: CreateKeyRequestBodyKeyUsage{
			value: "SIGN_VERIFY",
		},
	}
}

func (c CreateKeyRequestBodyKeyUsage) Value() string {
	return c.value
}

func (c CreateKeyRequestBodyKeyUsage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateKeyRequestBodyKeyUsage) UnmarshalJSON(b []byte) error {
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

type CreateKeyRequestBodyOrigin struct {
	value string
}

type CreateKeyRequestBodyOriginEnum struct {
	KMS      CreateKeyRequestBodyOrigin
	EXTERNAL CreateKeyRequestBodyOrigin
}

func GetCreateKeyRequestBodyOriginEnum() CreateKeyRequestBodyOriginEnum {
	return CreateKeyRequestBodyOriginEnum{
		KMS: CreateKeyRequestBodyOrigin{
			value: "kms",
		},
		EXTERNAL: CreateKeyRequestBodyOrigin{
			value: "external",
		},
	}
}

func (c CreateKeyRequestBodyOrigin) Value() string {
	return c.value
}

func (c CreateKeyRequestBodyOrigin) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateKeyRequestBodyOrigin) UnmarshalJSON(b []byte) error {
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
