package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogExportPathResp 训练作业日志输出信息。
type LogExportPathResp struct {

	// **参数解释**：训练作业日志保存的OBS地址，如：“obs://example/path”。 **取值范围**：不涉及。
	ObsUrl *string `json:"obs_url,omitempty"`

	// **参数解释**：训练作业日志保存的宿主机的路径，如：“/example/path”。 **取值范围**：不涉及。
	HostPath *string `json:"host_path,omitempty"`
}

func (o LogExportPathResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogExportPathResp struct{}"
	}

	return strings.Join([]string{"LogExportPathResp", string(data)}, " ")
}
