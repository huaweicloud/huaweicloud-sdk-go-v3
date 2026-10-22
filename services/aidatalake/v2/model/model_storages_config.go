package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StoragesConfig **参数解释**：存储资源配置（查询返回）。
type StoragesConfig struct {

	// **参数解释**：文件存储配置。 **约束限制**：最多10个文件存储配置。
	Files []FileStorageConfig `json:"files"`
}

func (o StoragesConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StoragesConfig struct{}"
	}

	return strings.Join([]string{"StoragesConfig", string(data)}, " ")
}
