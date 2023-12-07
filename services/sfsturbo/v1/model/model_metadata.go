package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Metadata metadata参数详情
type Metadata struct {

	// 要创加密文件系统，该字段传KMS服务专业版密钥的ID。
	CryptKeyId *string `json:"crypt_key_id,omitempty"`

	// 创专属文件系统，要创建的虚拟机的规格。
	DedicatedFlavor *string `json:"dedicated_flavor,omitempty"`

	// 创专属文件系统，要指定一个专属分布式存储的ID。
	DedicatedStorageId *string `json:"dedicated_storage_id,omitempty"`

	// 扩展类型。创建增强型/HPC型/HPC缓存型文件系统时，该参数必填。 创建增强型的文件系统，包括标准型-增强版和性能型-增强版，需要填写\"bandwidth\"。 创建HPC型文件系统，需要填写\"hpc\"。 创建HPC缓存型，需要填写\"hpc_cache\"。
	ExpandType *string `json:"expand_type,omitempty"`

	// 文件系统的带宽规格。创建HPC型/HPC缓存型文件系统时，该参数必填。 HPC型，可以填写\"20M\"、\"40M\"、\"125M\"、\"250M\"、\"500M\"、\"1000M\"。 HPC缓存型，可以填写\"2G\"、\"4G\"、\"8G\"、\"16G\"、\"24G\"、\"32G\"、\"48G\"。
	HpcBw *string `json:"hpc_bw,omitempty"`
}

func (o Metadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Metadata struct{}"
	}

	return strings.Join([]string{"Metadata", string(data)}, " ")
}
