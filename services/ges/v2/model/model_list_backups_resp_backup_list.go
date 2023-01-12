package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListBackupsRespBackupList struct {

	// 备份ID。
	Id *string `json:"id,omitempty"`

	// 备份名称。
	Name *string `json:"name,omitempty"`

	// 备份方法，取值为auto或manual。
	BackupMethod *string `json:"backup_method,omitempty"`

	// 备份关联的图ID。
	GraphId *string `json:"graph_id,omitempty"`

	// 备份关联的图Name。
	GraphName *string `json:"graph_name,omitempty"`

	// 备份关联的图状态。
	GraphStatus *string `json:"graph_status,omitempty"`

	// 备份关联的图规格。
	GraphSizeTypeIndex *string `json:"graph_size_type_index,omitempty"`

	// 备份关联的图版本。
	DataStoreVersion *string `json:"data_store_version,omitempty"`

	// 备份关联的图CPU架构。
	Arch *string `json:"arch,omitempty"`

	// 备份状态。  - backing_up：备份中 - success：备份成功 - failed：备份失败
	Status *string `json:"status,omitempty"`

	// 备份开始时间戳。
	StartTimestamp *int64 `json:"start_timestamp,omitempty"`

	// 备份开始时间。
	StartTime *string `json:"start_time,omitempty"`

	// 备份结束时间戳。
	EndTimestamp *int64 `json:"end_timestamp,omitempty"`

	// 备份结束时间。
	EndTime *string `json:"end_time,omitempty"`

	// 备份文件大小，单位为MB。
	Size *int64 `json:"size,omitempty"`

	// 备份时间，单位为秒。
	Duration *int64 `json:"duration,omitempty"`

	// 是否加密。true表示加密，默认值为\"false\"，不加密。
	Encrypted *bool `json:"encrypted,omitempty"`
}

func (o ListBackupsRespBackupList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupsRespBackupList struct{}"
	}

	return strings.Join([]string{"ListBackupsRespBackupList", string(data)}, " ")
}
