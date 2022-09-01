package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MultiTaskInitBody struct {
	ExtInfo *MultiTaskInitBodyExtInfo `json:"ext_info,omitempty" xml:"ext_info"`

	// 源端数据源ID
	SourceDatasourceId *string `json:"source_datasource_id,omitempty" xml:"source_datasource_id"`

	// 目标端数据源ID
	TargetDatasourceId *string `json:"target_datasource_id,omitempty" xml:"target_datasource_id"`

	// 任务ID，可以为空，为空时自动分配任务ID
	TaskId *string `json:"task_id,omitempty" xml:"task_id"`

	// 是否自动建立源端到目标端映射
	AutoMapping *bool `json:"auto_mapping,omitempty" xml:"auto_mapping"`
}

func (o MultiTaskInitBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiTaskInitBody struct{}"
	}

	return strings.Join([]string{"MultiTaskInitBody", string(data)}, " ")
}
