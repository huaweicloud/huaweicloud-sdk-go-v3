package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SummaryResp 可视化日志summary。
type SummaryResp struct {

	// **参数解释**：训练作业可视化日志类型，配置后训练作业可作为可视化作业数据源。 **取值范围**： - tensorboard：输出TensorBoard可视化工具类型的日志 - mindstudio-insight：输出mindstudio-insight可视化工具类型的日志
	LogType *string `json:"log_type,omitempty"`

	LogDir *LogDirResp `json:"log_dir,omitempty"`

	// **参数解释**：可视化作业或训练作业调试模式的可视化日志输入。
	DataSources *[]DataSourceResp `json:"data_sources,omitempty"`
}

func (o SummaryResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SummaryResp struct{}"
	}

	return strings.Join([]string{"SummaryResp", string(data)}, " ")
}
