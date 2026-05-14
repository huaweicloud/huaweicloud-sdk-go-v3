package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StartSyncTaskReq 启动同步任务body体
type StartSyncTaskReq struct {

	// 源端节点AK（最大长度100个字符）。URL列表迁移任务不需要填写此参数。
	SrcAk string `json:"src_ak"`

	// 源端节点SK（最大长度100个字符）。URL列表迁移任务不需要填写此参数。
	SrcSk string `json:"src_sk"`

	// 加解密类型，默认为DEFAULT，可选类型为DEFAULT、KMS
	SrcCryptoType *StartSyncTaskReqSrcCryptoType `json:"src_crypto_type,omitempty"`

	// KMS密钥ID，36个字符
	SrcKmsKeyId *string `json:"src_kms_key_id,omitempty"`

	// 目的端节点AK（最大长度100个字符）。
	DstAk string `json:"dst_ak"`

	// 目的端节点SK（最大长度100个字符）。
	DstSk string `json:"dst_sk"`

	// 加解密类型，默认为DEFAULT，可选类型为DEFAULT、KMS
	DstCryptoType *StartSyncTaskReqDstCryptoType `json:"dst_crypto_type,omitempty"`

	// KMS密钥ID，36个字符
	DstKmsKeyId *string `json:"dst_kms_key_id,omitempty"`

	// CDN鉴权密钥。
	SourceCdnAuthenticationKey *string `json:"source_cdn_authentication_key,omitempty"`

	// 加解密类型，默认为DEFAULT，可选类型为DEFAULT、KMS
	SourceCdnCryptoType *StartSyncTaskReqSourceCdnCryptoType `json:"source_cdn_crypto_type,omitempty"`

	// KMS密钥ID，36个字符
	SourceCdnKmsKeyId *string `json:"source_cdn_kms_key_id,omitempty"`
}

func (o StartSyncTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartSyncTaskReq struct{}"
	}

	return strings.Join([]string{"StartSyncTaskReq", string(data)}, " ")
}

type StartSyncTaskReqSrcCryptoType struct {
	value string
}

type StartSyncTaskReqSrcCryptoTypeEnum struct {
	DEFAULT StartSyncTaskReqSrcCryptoType
	KMS     StartSyncTaskReqSrcCryptoType
}

func GetStartSyncTaskReqSrcCryptoTypeEnum() StartSyncTaskReqSrcCryptoTypeEnum {
	return StartSyncTaskReqSrcCryptoTypeEnum{
		DEFAULT: StartSyncTaskReqSrcCryptoType{
			value: "DEFAULT",
		},
		KMS: StartSyncTaskReqSrcCryptoType{
			value: "KMS",
		},
	}
}

func (c StartSyncTaskReqSrcCryptoType) Value() string {
	return c.value
}

func (c StartSyncTaskReqSrcCryptoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartSyncTaskReqSrcCryptoType) UnmarshalJSON(b []byte) error {
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

type StartSyncTaskReqDstCryptoType struct {
	value string
}

type StartSyncTaskReqDstCryptoTypeEnum struct {
	DEFAULT StartSyncTaskReqDstCryptoType
	KMS     StartSyncTaskReqDstCryptoType
}

func GetStartSyncTaskReqDstCryptoTypeEnum() StartSyncTaskReqDstCryptoTypeEnum {
	return StartSyncTaskReqDstCryptoTypeEnum{
		DEFAULT: StartSyncTaskReqDstCryptoType{
			value: "DEFAULT",
		},
		KMS: StartSyncTaskReqDstCryptoType{
			value: "KMS",
		},
	}
}

func (c StartSyncTaskReqDstCryptoType) Value() string {
	return c.value
}

func (c StartSyncTaskReqDstCryptoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartSyncTaskReqDstCryptoType) UnmarshalJSON(b []byte) error {
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

type StartSyncTaskReqSourceCdnCryptoType struct {
	value string
}

type StartSyncTaskReqSourceCdnCryptoTypeEnum struct {
	DEFAULT StartSyncTaskReqSourceCdnCryptoType
	KMS     StartSyncTaskReqSourceCdnCryptoType
}

func GetStartSyncTaskReqSourceCdnCryptoTypeEnum() StartSyncTaskReqSourceCdnCryptoTypeEnum {
	return StartSyncTaskReqSourceCdnCryptoTypeEnum{
		DEFAULT: StartSyncTaskReqSourceCdnCryptoType{
			value: "DEFAULT",
		},
		KMS: StartSyncTaskReqSourceCdnCryptoType{
			value: "KMS",
		},
	}
}

func (c StartSyncTaskReqSourceCdnCryptoType) Value() string {
	return c.value
}

func (c StartSyncTaskReqSourceCdnCryptoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartSyncTaskReqSourceCdnCryptoType) UnmarshalJSON(b []byte) error {
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
