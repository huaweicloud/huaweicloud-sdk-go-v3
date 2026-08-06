package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskTableReferenceDetailResponse struct {

	// 作业id。
	JobId *int64 `json:"job_id,omitempty"`

	// 作业名。
	JobName *string `json:"job_name,omitempty"`

	// 数据库类型。
	DbType *string `json:"db_type,omitempty"`

	// 数据库名。
	DataBase *string `json:"data_base,omitempty"`

	// 数据表名。
	TableName *string `json:"table_name,omitempty"`

	// 集群名称。
	ClusterName *string `json:"cluster_name,omitempty"`

	// 作业的工作空间名。
	WorkspaceName *string `json:"workspace_name,omitempty"`

	// 作业的工作空间Id。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 作业责任人。
	Owner *string `json:"owner,omitempty"`

	// 作业最后提交时间。
	LastSubmitTime *int64 `json:"last_submit_time,omitempty"`

	// 作业和表的关系，0表示作业是读表，1表示作业写表。
	IoType *int32 `json:"io_type,omitempty"`

	// 是否是动态表。
	IsDynamic *bool `json:"is_dynamic,omitempty"`

	// 作业执行用户。
	ExecuteUser *string `json:"execute_user,omitempty"`
}

func (o TaskTableReferenceDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskTableReferenceDetailResponse struct{}"
	}

	return strings.Join([]string{"TaskTableReferenceDetailResponse", string(data)}, " ")
}
