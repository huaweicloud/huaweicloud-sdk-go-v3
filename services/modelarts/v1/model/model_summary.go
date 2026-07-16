package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Summary 可视化日志summary。
type Summary struct {

	// **参数解释**：训练作业可视化日志类型，配置后训练作业可作为可视化作业数据源。 **约束限制**：不涉及。 **取值范围**： - tensorboard：输出TensorBoard可视化工具类型的日志 - mindstudio-insight：输出mindstudio-insight可视化工具类型的日志  **默认取值**：不涉及。
	LogType *string `json:"log_type,omitempty"`

	LogDir *LogDir `json:"log_dir,omitempty"`

	// **参数解释**：可视化作业或训练作业调试模式的可视化日志输入。 **约束限制**：训练作业高级功能开启\"tensorboard/enable\": \"true\"或\"mindstudio-insight/enable\": \"true\"时必填。
	DataSources *[]DataSource `json:"data_sources,omitempty"`
}

func (o Summary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Summary struct{}"
	}

	return strings.Join([]string{"Summary", string(data)}, " ")
}
