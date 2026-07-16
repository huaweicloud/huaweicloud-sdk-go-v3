package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogExportPath 训练作业日志输出信息。
type LogExportPath struct {

	// 训练作业日志保存的OBS地址，如：“obs://example/path”。
	ObsUrl *string `json:"obs_url,omitempty"`

	// 训练作业日志保存的宿主机的路径，如：“/example/path”。
	HostPath *string `json:"host_path,omitempty"`
}

func (o LogExportPath) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogExportPath struct{}"
	}

	return strings.Join([]string{"LogExportPath", string(data)}, " ")
}
