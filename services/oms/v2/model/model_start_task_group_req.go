package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StartTaskGroupReq This is a auto create Body Object
type StartTaskGroupReq struct {

	// 源端节点AK（最大长度100个字符）。URL列表迁移任务不需要填写此参数。
	SrcAk *string `json:"src_ak,omitempty"`

	// 源端节点SK（最大长度100个字符）。URL列表迁移任务不需要填写此参数。
	SrcSk *string `json:"src_sk,omitempty"`

	// 加解密类型，默认为DEFAULT，可选类型为DEFAULT、KMS
	SrcCryptoType *StartTaskGroupReqSrcCryptoType `json:"src_crypto_type,omitempty"`

	// KMS密钥ID，36个字符
	SrcKmsKeyId *string `json:"src_kms_key_id,omitempty"`

	// 连接字符串，用于微软云Blob鉴权
	ConnectionString *string `json:"connection_string,omitempty"`

	// 用于谷歌云Cloud Storage鉴权
	JsonAuthFile *string `json:"json_auth_file,omitempty"`

	// 目的端节点AK（最大长度100个字符）。
	DstAk string `json:"dst_ak"`

	// 目的端节点SK（最大长度100个字符）。
	DstSk string `json:"dst_sk"`

	// 加解密类型，默认为DEFAULT，可选类型为DEFAULT、KMS
	DstCryptoType *StartTaskGroupReqDstCryptoType `json:"dst_crypto_type,omitempty"`

	// KMS密钥ID，36个字符
	DstKmsKeyId *string `json:"dst_kms_key_id,omitempty"`

	// CDN鉴权密钥。
	SourceCdnAuthenticationKey *string `json:"source_cdn_authentication_key,omitempty"`

	// 加解密类型，默认为DEFAULT，可选类型为DEFAULT、KMS
	SourceCdnCryptoType *StartTaskGroupReqSourceCdnCryptoType `json:"source_cdn_crypto_type,omitempty"`

	// KMS密钥ID，36个字符
	SourceCdnKmsKeyId *string `json:"source_cdn_kms_key_id,omitempty"`
}

func (o StartTaskGroupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartTaskGroupReq struct{}"
	}

	return strings.Join([]string{"StartTaskGroupReq", string(data)}, " ")
}

type StartTaskGroupReqSrcCryptoType struct {
	value string
}

type StartTaskGroupReqSrcCryptoTypeEnum struct {
	DEFAULT StartTaskGroupReqSrcCryptoType
	KMS     StartTaskGroupReqSrcCryptoType
}

func GetStartTaskGroupReqSrcCryptoTypeEnum() StartTaskGroupReqSrcCryptoTypeEnum {
	return StartTaskGroupReqSrcCryptoTypeEnum{
		DEFAULT: StartTaskGroupReqSrcCryptoType{
			value: "DEFAULT",
		},
		KMS: StartTaskGroupReqSrcCryptoType{
			value: "KMS",
		},
	}
}

func (c StartTaskGroupReqSrcCryptoType) Value() string {
	return c.value
}

func (c StartTaskGroupReqSrcCryptoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartTaskGroupReqSrcCryptoType) UnmarshalJSON(b []byte) error {
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

type StartTaskGroupReqDstCryptoType struct {
	value string
}

type StartTaskGroupReqDstCryptoTypeEnum struct {
	DEFAULT StartTaskGroupReqDstCryptoType
	KMS     StartTaskGroupReqDstCryptoType
}

func GetStartTaskGroupReqDstCryptoTypeEnum() StartTaskGroupReqDstCryptoTypeEnum {
	return StartTaskGroupReqDstCryptoTypeEnum{
		DEFAULT: StartTaskGroupReqDstCryptoType{
			value: "DEFAULT",
		},
		KMS: StartTaskGroupReqDstCryptoType{
			value: "KMS",
		},
	}
}

func (c StartTaskGroupReqDstCryptoType) Value() string {
	return c.value
}

func (c StartTaskGroupReqDstCryptoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartTaskGroupReqDstCryptoType) UnmarshalJSON(b []byte) error {
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

type StartTaskGroupReqSourceCdnCryptoType struct {
	value string
}

type StartTaskGroupReqSourceCdnCryptoTypeEnum struct {
	DEFAULT StartTaskGroupReqSourceCdnCryptoType
	KMS     StartTaskGroupReqSourceCdnCryptoType
}

func GetStartTaskGroupReqSourceCdnCryptoTypeEnum() StartTaskGroupReqSourceCdnCryptoTypeEnum {
	return StartTaskGroupReqSourceCdnCryptoTypeEnum{
		DEFAULT: StartTaskGroupReqSourceCdnCryptoType{
			value: "DEFAULT",
		},
		KMS: StartTaskGroupReqSourceCdnCryptoType{
			value: "KMS",
		},
	}
}

func (c StartTaskGroupReqSourceCdnCryptoType) Value() string {
	return c.value
}

func (c StartTaskGroupReqSourceCdnCryptoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartTaskGroupReqSourceCdnCryptoType) UnmarshalJSON(b []byte) error {
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
