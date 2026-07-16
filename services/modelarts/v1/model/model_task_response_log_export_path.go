package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskResponseLogExportPath **参数解释**：训练作业日志保存信息。
type TaskResponseLogExportPath struct {

	// **参数解释**：训练作业日志保存OBS路径。
	ObsUrl *string `json:"obs_url,omitempty"`
}

func (o TaskResponseLogExportPath) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskResponseLogExportPath struct{}"
	}

	return strings.Join([]string{"TaskResponseLogExportPath", string(data)}, " ")
}
