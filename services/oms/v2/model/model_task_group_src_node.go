package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TaskGroupSrcNode 迁移任务组的源端节点
type TaskGroupSrcNode struct {

	// 源端桶的AK（最大长度100个字符），task_type为非url_list时，本参数为必选。
	Ak *string `json:"ak,omitempty"`

	// 源端桶的SK（最大长度100个字符），task_type为非url_list时，本参数为必选。
	Sk *string `json:"sk,omitempty"`

	// 连接字符串，用于微软云Blob鉴权
	ConnectionString *string `json:"connection_string,omitempty"`

	// 加解密类型，默认为DEFAULT，可选类型为DEFAULT、KMS
	CryptoType *TaskGroupSrcNodeCryptoType `json:"crypto_type,omitempty"`

	// KMS密钥ID，36个字符
	KmsKeyId *string `json:"kms_key_id,omitempty"`

	// 用于谷歌云Cloud Storage鉴权
	JsonAuthFile *string `json:"json_auth_file,omitempty"`

	// 腾讯云APPID，当源端为腾讯云时，需要填写此参数，您可以在腾讯云控制台账号信息页面获取。
	AppId *string `json:"app_id,omitempty"`

	// 源端桶所处的区域，task_type为非URL_LIST时，本参数为必选。
	Region *string `json:"region,omitempty"`

	// 任务类型为前缀迁移任务时，表示待迁移前缀。 整桶迁移时，此参数设置为[\"\"]。
	ObjectKey *[]string `json:"object_key,omitempty"`

	// 源端所在桶
	Bucket *string `json:"bucket,omitempty"`

	// 源端云服务提供商，当task_type为URL_LIST时，本参数为URLSource且必选。可选值有AWS、Azure、Aliyun、Tencent、HuaweiCloud、QingCloud、KingsoftCloud、Baidu、Qiniu、URLSource、Google或者UCloud。默认值为Aliyun。
	CloudType *string `json:"cloud_type,omitempty"`

	ListFile *ListFile `json:"list_file,omitempty"`
}

func (o TaskGroupSrcNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskGroupSrcNode struct{}"
	}

	return strings.Join([]string{"TaskGroupSrcNode", string(data)}, " ")
}

type TaskGroupSrcNodeCryptoType struct {
	value string
}

type TaskGroupSrcNodeCryptoTypeEnum struct {
	DEFAULT TaskGroupSrcNodeCryptoType
	KMS     TaskGroupSrcNodeCryptoType
}

func GetTaskGroupSrcNodeCryptoTypeEnum() TaskGroupSrcNodeCryptoTypeEnum {
	return TaskGroupSrcNodeCryptoTypeEnum{
		DEFAULT: TaskGroupSrcNodeCryptoType{
			value: "DEFAULT",
		},
		KMS: TaskGroupSrcNodeCryptoType{
			value: "KMS",
		},
	}
}

func (c TaskGroupSrcNodeCryptoType) Value() string {
	return c.value
}

func (c TaskGroupSrcNodeCryptoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskGroupSrcNodeCryptoType) UnmarshalJSON(b []byte) error {
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
