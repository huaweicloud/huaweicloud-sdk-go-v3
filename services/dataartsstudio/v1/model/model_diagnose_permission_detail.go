package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DiagnosePermissionDetail struct {

	// 权限配置编号。
	Id *string `json:"id,omitempty"`

	// 诊断任务id。
	TaskId *string `json:"task_id,omitempty"`

	// 用户名。
	UserName *string `json:"user_name,omitempty"`

	// 权限类型。
	PermissionType *string `json:"permission_type,omitempty"`

	// 权限操作。
	PermissionAction *string `json:"permission_action,omitempty"`

	// 权限来源。
	PermissionSource *string `json:"permission_source,omitempty"`

	// 数据源类型。
	DatasourceType *string `json:"datasource_type,omitempty"`

	// 集群名称。
	ClusterName *string `json:"cluster_name,omitempty"`

	// 数据库名。
	Database *string `json:"database,omitempty"`

	// schema名。
	Schema *string `json:"schema,omitempty"`

	// 表名。
	Table *string `json:"table,omitempty"`

	// 备注。
	Remark *string `json:"remark,omitempty"`
}

func (o DiagnosePermissionDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiagnosePermissionDetail struct{}"
	}

	return strings.Join([]string{"DiagnosePermissionDetail", string(data)}, " ")
}
