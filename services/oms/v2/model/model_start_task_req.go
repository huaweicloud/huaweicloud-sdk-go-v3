package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StartTaskReq This is a auto create Body Object
type StartTaskReq struct {

	// 源端节点AK（最大长度100个字符）。URL列表迁移任务不需要填写此参数。
	SrcAk *string `json:"src_ak,omitempty"`

	// 源端节点SK（最大长度100个字符）。URL列表迁移任务不需要填写此参数。
	SrcSk *string `json:"src_sk,omitempty"`

	// 连接字符串，用于微软云Blob鉴权
	ConnectionString *string `json:"connection_string,omitempty"`

	// 加解密类型，默认为DEFAULT，可选类型为DEFAULT、KMS
	SrcCryptoType *StartTaskReqSrcCryptoType `json:"src_crypto_type,omitempty"`

	// KMS密钥ID，36个字符
	SrcKmsKeyId *string `json:"src_kms_key_id,omitempty"`

	// 用于谷歌云Cloud Storage鉴权
	JsonAuthFile *string `json:"json_auth_file,omitempty"`

	// 源端节点临时Token
	SrcSecurityToken *string `json:"src_security_token,omitempty"`

	// 目的端节点AK（最大长度100个字符）。
	DstAk string `json:"dst_ak"`

	// 目的端节点SK（最大长度100个字符）。
	DstSk string `json:"dst_sk"`

	// 加解密类型，默认为DEFAULT，可选类型为DEFAULT、KMS
	DstCryptoType *StartTaskReqDstCryptoType `json:"dst_crypto_type,omitempty"`

	// KMS密钥ID，36个字符
	DstKmsKeyId *string `json:"dst_kms_key_id,omitempty"`

	// 目标端节点临时Token
	DstSecurityToken *string `json:"dst_security_token,omitempty"`

	// CDN鉴权密钥。
	SourceCdnAuthenticationKey *string `json:"source_cdn_authentication_key,omitempty"`

	// 加解密类型，默认为DEFAULT，可选类型为DEFAULT、KMS
	SourceCdnCryptoType *StartTaskReqSourceCdnCryptoType `json:"source_cdn_crypto_type,omitempty"`

	// KMS密钥ID，36个字符
	SourceCdnKmsKeyId *string `json:"source_cdn_kms_key_id,omitempty"`

	// 迁移类型，标识是否为全量迁移，默认false（全量迁移）。 值为true时表示只重传失败对象。 值为空或者为false时表示全量迁移。
	MigrateFailedObject *bool `json:"migrate_failed_object,omitempty"`
}

func (o StartTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartTaskReq struct{}"
	}

	return strings.Join([]string{"StartTaskReq", string(data)}, " ")
}

type StartTaskReqSrcCryptoType struct {
	value string
}

type StartTaskReqSrcCryptoTypeEnum struct {
	DEFAULT StartTaskReqSrcCryptoType
	KMS     StartTaskReqSrcCryptoType
}

func GetStartTaskReqSrcCryptoTypeEnum() StartTaskReqSrcCryptoTypeEnum {
	return StartTaskReqSrcCryptoTypeEnum{
		DEFAULT: StartTaskReqSrcCryptoType{
			value: "DEFAULT",
		},
		KMS: StartTaskReqSrcCryptoType{
			value: "KMS",
		},
	}
}

func (c StartTaskReqSrcCryptoType) Value() string {
	return c.value
}

func (c StartTaskReqSrcCryptoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartTaskReqSrcCryptoType) UnmarshalJSON(b []byte) error {
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

type StartTaskReqDstCryptoType struct {
	value string
}

type StartTaskReqDstCryptoTypeEnum struct {
	DEFAULT StartTaskReqDstCryptoType
	KMS     StartTaskReqDstCryptoType
}

func GetStartTaskReqDstCryptoTypeEnum() StartTaskReqDstCryptoTypeEnum {
	return StartTaskReqDstCryptoTypeEnum{
		DEFAULT: StartTaskReqDstCryptoType{
			value: "DEFAULT",
		},
		KMS: StartTaskReqDstCryptoType{
			value: "KMS",
		},
	}
}

func (c StartTaskReqDstCryptoType) Value() string {
	return c.value
}

func (c StartTaskReqDstCryptoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartTaskReqDstCryptoType) UnmarshalJSON(b []byte) error {
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

type StartTaskReqSourceCdnCryptoType struct {
	value string
}

type StartTaskReqSourceCdnCryptoTypeEnum struct {
	DEFAULT StartTaskReqSourceCdnCryptoType
	KMS     StartTaskReqSourceCdnCryptoType
}

func GetStartTaskReqSourceCdnCryptoTypeEnum() StartTaskReqSourceCdnCryptoTypeEnum {
	return StartTaskReqSourceCdnCryptoTypeEnum{
		DEFAULT: StartTaskReqSourceCdnCryptoType{
			value: "DEFAULT",
		},
		KMS: StartTaskReqSourceCdnCryptoType{
			value: "KMS",
		},
	}
}

func (c StartTaskReqSourceCdnCryptoType) Value() string {
	return c.value
}

func (c StartTaskReqSourceCdnCryptoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartTaskReqSourceCdnCryptoType) UnmarshalJSON(b []byte) error {
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
