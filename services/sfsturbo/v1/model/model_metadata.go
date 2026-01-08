package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

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

	// 扩展类型；当文件系统正在创建时，该字段不返回。  - 创建增强型、20MB/s/TiB、40MB/s/TiB、125MB/s/TiB、250MB/s/TiB、500MB/s/TiB、1000MB/s/TiB、HPC缓存型文件系统时，该参数必填。  - 创建增强型的文件系统，包括标准型-增强版和性能型-增强版，需要填写\"bandwidth\"。  - 创建20MB/s/TiB、40MB/s/TiB、125MB/s/TiB、250MB/s/TiB、500MB/s/TiB、1000MB/s/TiB文件系统，需要填写\"hpc\"。  - 创建HPC缓存型，需要填写\"hpc_cache\"。
	ExpandType *string `json:"expand_type,omitempty"`

	// 文件系统的带宽规格。  创建20MB/s/TiB、40MB/s/TiB、125MB/s/TiB、250MB/s/TiB、500MB/s/TiB、1000MB/s/TiB、HPC缓存型文件系统时，该参数必填。  20MB/s/TiB，填写\"20M\"。 40MB/s/TiB，填写\"40M\"。 125MB/s/TiB，填写\"125M\"。 250MB/s/TiB，填写\"250M\"。 500MB/s/TiB，填写\"500M\"。 1000MB/s/TiB，填写\"1000M\"。 HPC缓存型，填写\"2G\"、\"4G\"、\"8G\"、\"16G\"、\"24G\"、\"32G\"或\"48G\"。
	HpcBw *string `json:"hpc_bw,omitempty"`

	// 是否自动创建安全组规则。\"true\"表示自动创建安全组规则，\"false\"表示不自动创建安全组规则。默认值是\"true\"。
	AutoCreateSecurityGroupRules *MetadataAutoCreateSecurityGroupRules `json:"auto_create_security_group_rules,omitempty"`

	// 存储库ID。
	VaultId *string `json:"vault_id,omitempty"`
}

func (o Metadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Metadata struct{}"
	}

	return strings.Join([]string{"Metadata", string(data)}, " ")
}

type MetadataAutoCreateSecurityGroupRules struct {
	value string
}

type MetadataAutoCreateSecurityGroupRulesEnum struct {
	TRUE  MetadataAutoCreateSecurityGroupRules
	FALSE MetadataAutoCreateSecurityGroupRules
}

func GetMetadataAutoCreateSecurityGroupRulesEnum() MetadataAutoCreateSecurityGroupRulesEnum {
	return MetadataAutoCreateSecurityGroupRulesEnum{
		TRUE: MetadataAutoCreateSecurityGroupRules{
			value: "true",
		},
		FALSE: MetadataAutoCreateSecurityGroupRules{
			value: "false",
		},
	}
}

func (c MetadataAutoCreateSecurityGroupRules) Value() string {
	return c.value
}

func (c MetadataAutoCreateSecurityGroupRules) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MetadataAutoCreateSecurityGroupRules) UnmarshalJSON(b []byte) error {
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
