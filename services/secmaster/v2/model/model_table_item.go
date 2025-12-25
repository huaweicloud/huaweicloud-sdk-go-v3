package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TableItem struct {

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 工作空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// UUID
	TableId *string `json:"table_id,omitempty"`

	// Table name 表名
	TableName *string `json:"table_name,omitempty"`

	// 表别名
	TableAlias *string `json:"table_alias,omitempty"`

	// directory 目录分组
	Directory *string `json:"directory,omitempty"`

	// 表描述
	Description *string `json:"description,omitempty"`

	Category *TableCategory `json:"category,omitempty"`

	LockStatus *TableLockStatus `json:"lock_status,omitempty"`

	ProcessStatus *TableProcessStatus `json:"process_status,omitempty"`

	ProcessError *TableProcessError `json:"process_error,omitempty"`

	EditType *TableEditType `json:"edit_type,omitempty"`

	Format *TableFormat `json:"format,omitempty"`

	RwType *TableRwType `json:"rw_type,omitempty"`

	OwnerType *TableOwnerType `json:"owner_type,omitempty"`

	DataLayering *DataLayering `json:"data_layering,omitempty"`

	DataClassification *DataClassification `json:"data_classification,omitempty"`

	// 毫秒时间戳
	CreateTime *int64 `json:"create_time,omitempty"`

	// 毫秒时间戳
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 毫秒时间戳
	DeleteTime *int64 `json:"delete_time,omitempty"`

	// 已使用存储容量
	StoreSizeInBytes *int64 `json:"store_size_in_bytes,omitempty"`
}

func (o TableItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableItem struct{}"
	}

	return strings.Join([]string{"TableItem", string(data)}, " ")
}
