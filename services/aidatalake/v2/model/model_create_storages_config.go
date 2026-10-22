package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStoragesConfig **参数解释**：存储资源配置（创建时）。
type CreateStoragesConfig struct {

	// **参数解释**：文件存储配置。 **约束限制**：最多10个文件存储配置。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Files *[]CreateFileStorageConfig `json:"files,omitempty"`
}

func (o CreateStoragesConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStoragesConfig struct{}"
	}

	return strings.Join([]string{"CreateStoragesConfig", string(data)}, " ")
}
