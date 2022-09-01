package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MultiTaskInitElement struct {
	ExtInfo *MultiTaskInitElementExtInfo `json:"ext_info,omitempty" xml:"ext_info"`

	// 任务ID
	TaskId *string `json:"task_id,omitempty" xml:"task_id"`

	// 源端数据源ID
	SourceDatasourceId *string `json:"source_datasource_id,omitempty" xml:"source_datasource_id"`

	// 目标端数据源ID
	TargetDatasourceId *string `json:"target_datasource_id,omitempty" xml:"target_datasource_id"`

	// 源端组
	SourceGroup *string `json:"source_group,omitempty" xml:"source_group"`

	// 目标端组
	TargetGroup *string `json:"target_group,omitempty" xml:"target_group"`

	// 源端数据源ID
	SourceDsId *string `json:"source_ds_id,omitempty" xml:"source_ds_id"`

	// 目标端数据源ID
	TargetDsId *string `json:"target_ds_id,omitempty" xml:"target_ds_id"`

	// 源端实例ID
	SourceInstanceId *string `json:"source_instance_id,omitempty" xml:"source_instance_id"`

	// 目标端实例ID
	TargetInstanceId *string `json:"target_instance_id,omitempty" xml:"target_instance_id"`

	// 源端数据源所属集成应用ID
	SourceAppId *string `json:"source_app_id,omitempty" xml:"source_app_id"`

	// 目标端数据源所属集成应用ID
	TargetAppId *string `json:"target_app_id,omitempty" xml:"target_app_id"`

	// 源端数据源的名称
	SourceDatasourceName *string `json:"source_datasource_name,omitempty" xml:"source_datasource_name"`

	// 目标端数据源的名称
	TargetDatasourceName *string `json:"target_datasource_name,omitempty" xml:"target_datasource_name"`

	// 源端数据源的类型
	SourceDatasourceType *string `json:"source_datasource_type,omitempty" xml:"source_datasource_type"`

	// 目标端数据源的类型
	TargetDatasourceType *string `json:"target_datasource_type,omitempty" xml:"target_datasource_type"`

	// 映射关系列表，只返回前10条
	Mappings *[]MultiTaskMappingElement `json:"mappings,omitempty" xml:"mappings"`

	// 映射关系总数
	MappingsTotalCount *int64 `json:"mappings_total_count,omitempty" xml:"mappings_total_count"`
}

func (o MultiTaskInitElement) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiTaskInitElement struct{}"
	}

	return strings.Join([]string{"MultiTaskInitElement", string(data)}, " ")
}
