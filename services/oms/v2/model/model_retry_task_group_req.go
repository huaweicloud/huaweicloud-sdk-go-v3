package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RetryTaskGroupReq 重试迁移任务组请求参数
type RetryTaskGroupReq struct {

	// 源端ak（最大长度100个字符）
	SrcAk *string `json:"src_ak,omitempty"`

	// 源端sk（最大长度100个字符）
	SrcSk *string `json:"src_sk,omitempty"`

	// 加解密类型，默认为DEFAULT，可选类型为DEFAULT、KMS
	SrcCryptoType *RetryTaskGroupReqSrcCryptoType `json:"src_crypto_type,omitempty"`

	// KMS密钥ID，36个字符
	SrcKmsKeyId *string `json:"src_kms_key_id,omitempty"`

	// 连接字符串，用于微软云Blob鉴权
	ConnectionString *string `json:"connection_string,omitempty"`

	// 用于谷歌云Cloud Storage鉴权
	JsonAuthFile *string `json:"json_auth_file,omitempty"`

	// 目的端ak（最大长度100个字符）
	DstAk *string `json:"dst_ak,omitempty"`

	// 目的端sk（最大长度100个字符）
	DstSk *string `json:"dst_sk,omitempty"`

	// 加解密类型，默认为DEFAULT，可选类型为DEFAULT、KMS
	DstCryptoType *RetryTaskGroupReqDstCryptoType `json:"dst_crypto_type,omitempty"`

	// KMS密钥ID，36个字符
	DstKmsKeyId *string `json:"dst_kms_key_id,omitempty"`

	// cdn鉴权密钥
	SourceCdnAuthenticationKey *string `json:"source_cdn_authentication_key,omitempty"`

	// 加解密类型，默认为DEFAULT，可选类型为DEFAULT、KMS
	SourceCdnCryptoType *RetryTaskGroupReqSourceCdnCryptoType `json:"source_cdn_crypto_type,omitempty"`

	// KMS密钥ID，36个字符
	SourceCdnKmsKeyId *string `json:"source_cdn_kms_key_id,omitempty"`

	// 失败任务重试方式，标识是否为全量重新迁移，默认false（全量重新迁移）。 值为true时表示只重传失败对象。 值为空或者为false时表示全量重新迁移（默认跳过目的端已迁移对象）。
	MigrateFailedObject *bool `json:"migrate_failed_object,omitempty"`
}

func (o RetryTaskGroupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryTaskGroupReq struct{}"
	}

	return strings.Join([]string{"RetryTaskGroupReq", string(data)}, " ")
}

type RetryTaskGroupReqSrcCryptoType struct {
	value string
}

type RetryTaskGroupReqSrcCryptoTypeEnum struct {
	DEFAULT RetryTaskGroupReqSrcCryptoType
	KMS     RetryTaskGroupReqSrcCryptoType
}

func GetRetryTaskGroupReqSrcCryptoTypeEnum() RetryTaskGroupReqSrcCryptoTypeEnum {
	return RetryTaskGroupReqSrcCryptoTypeEnum{
		DEFAULT: RetryTaskGroupReqSrcCryptoType{
			value: "DEFAULT",
		},
		KMS: RetryTaskGroupReqSrcCryptoType{
			value: "KMS",
		},
	}
}

func (c RetryTaskGroupReqSrcCryptoType) Value() string {
	return c.value
}

func (c RetryTaskGroupReqSrcCryptoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RetryTaskGroupReqSrcCryptoType) UnmarshalJSON(b []byte) error {
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

type RetryTaskGroupReqDstCryptoType struct {
	value string
}

type RetryTaskGroupReqDstCryptoTypeEnum struct {
	DEFAULT RetryTaskGroupReqDstCryptoType
	KMS     RetryTaskGroupReqDstCryptoType
}

func GetRetryTaskGroupReqDstCryptoTypeEnum() RetryTaskGroupReqDstCryptoTypeEnum {
	return RetryTaskGroupReqDstCryptoTypeEnum{
		DEFAULT: RetryTaskGroupReqDstCryptoType{
			value: "DEFAULT",
		},
		KMS: RetryTaskGroupReqDstCryptoType{
			value: "KMS",
		},
	}
}

func (c RetryTaskGroupReqDstCryptoType) Value() string {
	return c.value
}

func (c RetryTaskGroupReqDstCryptoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RetryTaskGroupReqDstCryptoType) UnmarshalJSON(b []byte) error {
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

type RetryTaskGroupReqSourceCdnCryptoType struct {
	value string
}

type RetryTaskGroupReqSourceCdnCryptoTypeEnum struct {
	DEFAULT RetryTaskGroupReqSourceCdnCryptoType
	KMS     RetryTaskGroupReqSourceCdnCryptoType
}

func GetRetryTaskGroupReqSourceCdnCryptoTypeEnum() RetryTaskGroupReqSourceCdnCryptoTypeEnum {
	return RetryTaskGroupReqSourceCdnCryptoTypeEnum{
		DEFAULT: RetryTaskGroupReqSourceCdnCryptoType{
			value: "DEFAULT",
		},
		KMS: RetryTaskGroupReqSourceCdnCryptoType{
			value: "KMS",
		},
	}
}

func (c RetryTaskGroupReqSourceCdnCryptoType) Value() string {
	return c.value
}

func (c RetryTaskGroupReqSourceCdnCryptoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RetryTaskGroupReqSourceCdnCryptoType) UnmarshalJSON(b []byte) error {
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
