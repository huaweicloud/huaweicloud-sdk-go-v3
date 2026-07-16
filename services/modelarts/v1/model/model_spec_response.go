package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SpecResponse 训练作业规格参数。
type SpecResponse struct {
	Resource *Resource `json:"resource,omitempty"`

	// **参数解释**：训练作业挂载卷信息。
	Volumes *[]JobVolumeResp `json:"volumes,omitempty"`

	LogExportPath *LogExportPathResp `json:"log_export_path,omitempty"`

	SchedulePolicy *SchedulePolicyResp `json:"schedule_policy,omitempty"`

	// **参数解释**：指标采集配置。
	CustomMetrics *[]CustomMetrics `json:"custom_metrics,omitempty"`

	OutputModel *OutputModelResp `json:"output_model,omitempty"`

	AssetModel *AssetModelResp `json:"asset_model,omitempty"`
}

func (o SpecResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpecResponse struct{}"
	}

	return strings.Join([]string{"SpecResponse", string(data)}, " ")
}
