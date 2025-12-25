package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTableResponse Response Object
type CreateTableResponse struct {

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 工作空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// UUID
	TableId *string `json:"table_id,omitempty"`

	// UUID
	PipeId *string `json:"pipe_id,omitempty"`

	// 表名称
	TableName *string `json:"table_name,omitempty"`

	// 表别名
	TableAlias *string `json:"table_alias,omitempty"`

	// 相关描述
	Description *string `json:"description,omitempty"`

	// directory 目录分组
	Directory *string `json:"directory,omitempty"`

	Category *IsapTableCategory `json:"category,omitempty"`

	LockStatus *TableLockStatus `json:"lock_status,omitempty"`

	ProcessStatus *IsapTableProcessStatus `json:"process_status,omitempty"`

	ProcessError *IsapTableProcessError `json:"process_error,omitempty"`

	Format *TableFormat `json:"format,omitempty"`

	RwType *TableRwType `json:"rw_type,omitempty"`

	OwnerType *TableOwnerType `json:"owner_type,omitempty"`

	DataLayering *DataLayering `json:"data_layering,omitempty"`

	DataClassification *DataClassification `json:"data_classification,omitempty"`

	Schema *IsapTableSchema `json:"schema,omitempty"`

	StorageSetting *TableStorageSetting `json:"storage_setting,omitempty"`

	DisplaySetting *TableDisplaySetting `json:"display_setting,omitempty"`

	// 毫秒时间戳
	CreateTime *int64 `json:"create_time,omitempty"`

	// 毫秒时间戳
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 毫秒时间戳
	DeleteTime     *int64 `json:"delete_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTableResponse struct{}"
	}

	return strings.Join([]string{"CreateTableResponse", string(data)}, " ")
}
