package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsapResource 资源
type IsapResource struct {

	// 资源类型，例如 STREAMING 或 INDEX
	Category *string `json:"category,omitempty"`

	// 创建者信息
	CreateBy *string `json:"create_by,omitempty"`

	// 创建时间戳
	CreateTime *int32 `json:"create_time,omitempty"`

	// 数据分类，例如 FACTUAL_DATA
	DataClassification *string `json:"data_classification,omitempty"`

	// 数据分层，例如 ODS
	DataLayering *string `json:"data_layering,omitempty"`

	// 删除时间戳
	DeleteTime *int32 `json:"delete_time,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 目录路径
	Directory *string `json:"directory,omitempty"`

	DisplaySetting *IsapResourceDisplaySetting `json:"display_setting,omitempty"`

	// 数据格式，例如 JSON
	Format *string `json:"format,omitempty"`

	// 锁定状态，例如 UNLOCKED
	LockStatus *string `json:"lock_status,omitempty"`

	// 所有者类型，例如 USER
	OwnerType *string `json:"owner_type,omitempty"`

	// 管道ID
	PipeId *string `json:"pipe_id,omitempty"`

	// 处理错误状态，例如 NONE
	ProcessError *string `json:"process_error,omitempty"`

	// 处理状态，例如 COMPLETED
	ProcessStatus *string `json:"process_status,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 读写类型，例如 READ_WRITE
	RwType *string `json:"rw_type,omitempty"`

	Schema *Schema `json:"schema,omitempty"`

	StorageSetting *StorageSetting `json:"storage_setting,omitempty"`

	// 存储大小（字节）
	StoreSizeInBytes *int32 `json:"store_size_in_bytes,omitempty"`

	// 表别名
	TableAlias *string `json:"table_alias,omitempty"`

	// 表ID
	TableId *string `json:"table_id,omitempty"`

	// 表名称
	TableName *string `json:"table_name,omitempty"`

	// 更新者信息
	UpdateBy *string `json:"update_by,omitempty"`

	// 更新时间戳
	UpdateTime *int32 `json:"update_time,omitempty"`

	// 工作空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`
}

func (o IsapResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapResource struct{}"
	}

	return strings.Join([]string{"IsapResource", string(data)}, " ")
}
