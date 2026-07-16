package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogExportConfig 日志导出配置
type LogExportConfig struct {

	// **参数解释**：日志版本号。 **约束限制**： - 日志版本取值为v0、v1，默认为v0。 **取值范围**：v0、v1 **默认取值**：v0。
	Version *string `json:"version,omitempty"`

	// **参数解释**：是否开启日志分时段下载。 **约束限制**：不涉及。 **取值范围**： - true：开启日志分时段下载 - false：关闭日志分时段下载 **默认取值**：false。
	RotationEnabled *bool `json:"rotation_enabled,omitempty"`
}

func (o LogExportConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogExportConfig struct{}"
	}

	return strings.Join([]string{"LogExportConfig", string(data)}, " ")
}
