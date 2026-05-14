package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SrcNodeReq 源端节点信息。
type SrcNodeReq struct {

	// 源端云服务提供商，task_type为非url_list时，本参数为URLSource且必选。  可选值有AWS、Azure、Aliyun、Tencent、HuaweiCloud、QingCloud、KingsoftCloud、Baidu、Qiniu、Google、URLSource或者UCloud。默认值为Aliyun。
	CloudType *string `json:"cloud_type,omitempty"`

	// 源端桶所处的区域，task_type为非url_list时，本参数为必选。
	Region *string `json:"region,omitempty"`

	// 源端桶的AK（最大长度100个字符），task_type为非url_list时，本参数为必选。
	Ak *string `json:"ak,omitempty"`

	// 源端桶的SK（最大长度100个字符），task_type为非url_list时，本参数为必选。
	Sk *string `json:"sk,omitempty"`

	// 连接字符串，用于微软云Blob鉴权
	ConnectionString *string `json:"connection_string,omitempty"`

	// 用于谷歌云Cloud Storage鉴权
	JsonAuthFile *string `json:"json_auth_file,omitempty"`

	// 源端桶的临时Token（最大长度16384个字符）
	SecurityToken *string `json:"security_token,omitempty"`

	// 腾讯云APPID，当源端为腾讯云时，需要填写此参数，您可以在腾讯云控制台账号信息页面获取。
	AppId *string `json:"app_id,omitempty"`

	// 源端桶的名称，task_type为非url_list时，本参数为必选。
	Bucket *string `json:"bucket,omitempty"`

	// 任务类型为对象迁移任务时，表示待迁移对象名称（以“/”结尾的字符串代表待迁移的文件夹，非“/”结尾的字符串代表待迁移的文件。）； 任务类型为前缀迁移任务时，表示待迁移前缀。 整桶迁移时，此参数设置为[\"\"]。
	ObjectKey *[]string `json:"object_key,omitempty"`

	ListFile *ListFile `json:"list_file,omitempty"`

	// 加解密类型，默认为DEFAULT，可选类型为DEFAULT、KMS
	CryptoType *SrcNodeReqCryptoType `json:"crypto_type,omitempty"`

	// KMS密钥ID，36个字符
	KmsKeyId *string `json:"kms_key_id,omitempty"`
}

func (o SrcNodeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SrcNodeReq struct{}"
	}

	return strings.Join([]string{"SrcNodeReq", string(data)}, " ")
}

type SrcNodeReqCryptoType struct {
	value string
}

type SrcNodeReqCryptoTypeEnum struct {
	DEFAULT SrcNodeReqCryptoType
	KMS     SrcNodeReqCryptoType
}

func GetSrcNodeReqCryptoTypeEnum() SrcNodeReqCryptoTypeEnum {
	return SrcNodeReqCryptoTypeEnum{
		DEFAULT: SrcNodeReqCryptoType{
			value: "DEFAULT",
		},
		KMS: SrcNodeReqCryptoType{
			value: "KMS",
		},
	}
}

func (c SrcNodeReqCryptoType) Value() string {
	return c.value
}

func (c SrcNodeReqCryptoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SrcNodeReqCryptoType) UnmarshalJSON(b []byte) error {
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
