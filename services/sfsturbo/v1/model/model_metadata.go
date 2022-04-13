package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// metadata参数详情
type Metadata struct {
	// 要创加密文件系统，该字段传KMS服务专业版密钥的ID。

	CryptKeyId *string `json:"crypt_key_id,omitempty"`
	// 创专属文件系统，要创建的虚拟机的规格。

	DedicatedFlavor *string `json:"dedicated_flavor,omitempty"`
	// 创专属文件系统，要指定一个专属分布式存储的ID。

	DedicatedStorageId *string `json:"dedicated_storage_id,omitempty"`
	// 扩展类型。当前有效值为bandwidth，即创建增强型的文件系统。

	ExpandType *string `json:"expand_type,omitempty"`
}

func (o Metadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Metadata struct{}"
	}

	return strings.Join([]string{"Metadata", string(data)}, " ")
}
