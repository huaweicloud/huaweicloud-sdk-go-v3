package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DiagnoseNoMaskingDetail 未脱敏的表详情
type DiagnoseNoMaskingDetail struct {

	// 详情uuid
	Id *string `json:"id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 诊断任务id
	TaskId *string `json:"task_id,omitempty"`

	// 数据源类型
	DatasourceType *string `json:"datasource_type,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 数据库名称
	Database *string `json:"database,omitempty"`

	// schema名称
	Schema *string `json:"schema,omitempty"`

	// 表名称
	Table *string `json:"table,omitempty"`

	// 详情评价
	Remark *string `json:"remark,omitempty"`
}

func (o DiagnoseNoMaskingDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiagnoseNoMaskingDetail struct{}"
	}

	return strings.Join([]string{"DiagnoseNoMaskingDetail", string(data)}, " ")
}
