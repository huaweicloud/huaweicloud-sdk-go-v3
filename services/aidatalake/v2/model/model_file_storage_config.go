package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FileStorageConfig 文件存储配置
type FileStorageConfig struct {

	// **参数解释**：存储资源ID，可通过[查询存储资源列表](ListV2Storages.xml)获取存储资源ID。 **约束限制**：不涉及。 **取值范围**：长度为1~36的英文字符、数字和中划线的组合。 **默认取值**：不涉及。
	StorageId string `json:"storage_id"`

	// **参数解释**：存储的名称，对应PVC的名称。 **约束限制**：不涉及。 **取值范围**：支持小写字母、数字、中划线的组合，长度为1~63个字符。 **默认取值**：不涉及。
	StorageName *string `json:"storage_name,omitempty"`

	// **参数解释**：存储上的路径。 **约束限制**：不涉及。 **取值范围**：支持英文字母、数字、斜杠及特殊字符，最小长度为1，最大长度为1024。 **默认取值**：不涉及。
	StoragePath *string `json:"storage_path,omitempty"`

	// **参数解释**：挂载到容器的路径。 **约束限制**：不涉及。 **取值范围**：支持英文字母、数字、斜杠及特殊字符，最小长度为1，最大长度为1024。 **默认取值**：不涉及。
	MountPath *string `json:"mount_path,omitempty"`
}

func (o FileStorageConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileStorageConfig struct{}"
	}

	return strings.Join([]string{"FileStorageConfig", string(data)}, " ")
}
