package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFileStorageConfig **参数解释**：文件存储配置。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
type CreateFileStorageConfig struct {

	// **参数解释**：存储资源ID，可通过[查询存储资源列表](ListV2Storages.xml)获取存储资源ID。 **约束限制**：不涉及。 **取值范围**：长度为1~36的英文字符、数字和中划线的组合。 **默认取值**：不涉及。
	StorageId string `json:"storage_id"`

	// **参数解释**：存储上的路径。 **约束限制**：存储支持的格式。 **取值范围**：最小长度为1，最大长度为1024。 **默认取值**：/。
	StoragePath *string `json:"storage_path,omitempty"`

	// **参数解释**：挂载到容器的路径。 **约束限制**：容器本地支持的格式，且不能覆盖系统路径，限制必须挂载在/home/service/下。 **取值范围**：最小长度为1，最大长度为1024。 **默认取值**：/home/service/works。
	MountPath *string `json:"mount_path,omitempty"`
}

func (o CreateFileStorageConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFileStorageConfig struct{}"
	}

	return strings.Join([]string{"CreateFileStorageConfig", string(data)}, " ")
}
