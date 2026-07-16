package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataSource 可视化作业或训练作业调试模式的可视化日志输入，训练作业高级功能开启\"tensorboard/enable\": \"true\"或\"mindstudio-insight/enable\": \"true\"时必填。
type DataSource struct {
	Job *JobSummary `json:"job"`

	Nfs *NfsSummary `json:"nfs,omitempty"`
}

func (o DataSource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataSource struct{}"
	}

	return strings.Join([]string{"DataSource", string(data)}, " ")
}
