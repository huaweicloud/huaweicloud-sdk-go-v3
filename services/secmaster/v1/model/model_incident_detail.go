package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IncidentDetail struct {

	// 记录时间，格式ISO8601：YYYY-MM-DDTHH:mm:ss.ms+timezone。时区信息为事件发生时区，无法解析时区的时间，默认时区填东八区
	CreateTime *string `json:"create_time,omitempty"`

	DataObject *Incident `json:"data_object,omitempty"`

	DataclassRef *AlertDetailDataclassRef `json:"dataclass_ref,omitempty"`

	// 格式版本
	FormatVersion *int32 `json:"format_version,omitempty"`

	// 事件唯一标识，UUID格式，最大36个字符
	Id *string `json:"id,omitempty"`

	// 当前项目的id
	ProjectId *string `json:"project_id,omitempty"`

	// 更新时间，格式ISO8601：YYYY-MM-DDTHH:mm:ss.ms+timezone。时区信息为告警发生时区，无法解析时区的时间，默认时区填东八区
	UpdateTime *string `json:"update_time,omitempty"`

	// 版本
	Version *int32 `json:"version,omitempty"`

	// 当前的工作空间id
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 数据类型（SIMULATION--告警, PLAYBOOK--剧本, MANUAL--手动触发,DATA_SOURCE--数据源）
	Type *string `json:"type,omitempty"`
}

func (o IncidentDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IncidentDetail struct{}"
	}

	return strings.Join([]string{"IncidentDetail", string(data)}, " ")
}
