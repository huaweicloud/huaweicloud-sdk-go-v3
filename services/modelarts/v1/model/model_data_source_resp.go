package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataSourceResp 可视化作业或训练作业调试模式的可视化日志输入，训练作业高级功能开启\"tensorboard/enable\": \"true\"或\"mindstudio-insight/enable\": \"true\"时必填。
type DataSourceResp struct {
	Job *JobSummaryResp `json:"job"`
}

func (o DataSourceResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataSourceResp struct{}"
	}

	return strings.Join([]string{"DataSourceResp", string(data)}, " ")
}
