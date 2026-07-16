package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskLogExportPath **参数解释**：训练作业日志保存信息。 **约束限制**：不涉及。
type TaskLogExportPath struct {

	// **参数解释**：训练作业日志保存OBS路径。 **约束限制**：不涉及。
	ObsUrl *string `json:"obs_url,omitempty"`
}

func (o TaskLogExportPath) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskLogExportPath struct{}"
	}

	return strings.Join([]string{"TaskLogExportPath", string(data)}, " ")
}
