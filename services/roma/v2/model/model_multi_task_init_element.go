package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MultiTaskInitElement struct {
	ExtInfo *MultiTaskInitElementExtInfo `json:"ext_info,omitempty"`
	// 任务ID

	TaskId *string `json:"task_id,omitempty"`
	// 源端数据源ID

	SourceDatasourceId *string `json:"source_datasource_id,omitempty"`
	// 目标端数据源ID

	TargetDatasourceId *string `json:"target_datasource_id,omitempty"`
	// 源端组

	SourceGroup *string `json:"source_group,omitempty"`
	// 目标端组

	TargetGroup *string `json:"target_group,omitempty"`
	// 源端数据源ID

	SourceDsId *string `json:"source_ds_id,omitempty"`
	// 目标端数据源ID

	TargetDsId *string `json:"target_ds_id,omitempty"`
	// 源端实例ID

	SourceInstanceId *string `json:"source_instance_id,omitempty"`
	// 目标端实例ID

	TargetInstanceId *string `json:"target_instance_id,omitempty"`
	// 源端数据源所属集成应用ID

	SourceAppId *string `json:"source_app_id,omitempty"`
	// 目标端数据源所属集成应用ID

	TargetAppId *string `json:"target_app_id,omitempty"`
	// 源端数据源的名称

	SourceDatasourceName *string `json:"source_datasource_name,omitempty"`
	// 目标端数据源的名称

	TargetDatasourceName *string `json:"target_datasource_name,omitempty"`
	// 源端数据源的类型

	SourceDatasourceType *string `json:"source_datasource_type,omitempty"`
	// 目标端数据源的类型

	TargetDatasourceType *string `json:"target_datasource_type,omitempty"`
	// 映射关系列表，只返回前10条

	Mappings *[]MultiTaskMappingElement `json:"mappings,omitempty"`
	// 映射关系总数

	MappingsTotalCount *int64 `json:"mappings_total_count,omitempty"`
}

func (o MultiTaskInitElement) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiTaskInitElement struct{}"
	}

	return strings.Join([]string{"MultiTaskInitElement", string(data)}, " ")
}
