package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Spec 训练作业规格参数。有此字段时，无需填写tasks字段。
type Spec struct {
	Resource *SpecResource `json:"resource,omitempty"`

	// **参数解释**：训练作业挂载卷信息。 **约束限制**：不涉及。
	Volumes *[]SpecVolumes `json:"volumes,omitempty"`

	LogExportPath *LogExportPath `json:"log_export_path,omitempty"`

	AutoStop *AutoStop `json:"auto_stop,omitempty"`

	SchedulePolicy *SchedulePolicy `json:"schedule_policy,omitempty"`

	LogExportConfig *LogExportConfig `json:"log_export_config,omitempty"`

	Notification *Notification `json:"notification,omitempty"`

	// **参数解释**：指标采集配置。
	CustomMetrics *[]CustomMetrics `json:"custom_metrics,omitempty"`

	OutputModel *OutputModel `json:"output_model,omitempty"`

	AssetModel *AssetModel `json:"asset_model,omitempty"`

	// **参数解释**：精调训练作业资产模型ID。
	AssetId *string `json:"asset_id,omitempty"`
}

func (o Spec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Spec struct{}"
	}

	return strings.Join([]string{"Spec", string(data)}, " ")
}
