package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Backup 图备份
type Backup struct {

	// 备份ID。
	Id string `json:"id"`

	// 备份名称。
	Name string `json:"name"`

	// 备份方法，取值为auto或manual。
	BackupMethod string `json:"backupMethod"`

	// 备份关联的图ID。
	GraphId string `json:"graphId"`

	// 备份关联的图Name。
	GraphName string `json:"graphName"`

	// 备份关联的图状态。
	GraphStatus string `json:"graphStatus"`

	// 备份关联的图规格。
	GraphSizeTypeIndex string `json:"graphSizeTypeIndex"`

	// 备份关联的图版本。
	DataStoreVersion string `json:"dataStoreVersion"`

	// 备份关联的图CPU架构。
	Arch string `json:"arch"`

	// 备份状态。  - backing_up：备份中 - success：备份成功 - failed：备份失败
	Status string `json:"status"`

	// 备份开始时间戳。
	StartTimestamp int64 `json:"startTimestamp"`

	// 备份时间。
	StartTime string `json:"startTime"`

	// 备份结束时间戳。
	EndTimestamp int64 `json:"endTimestamp"`

	// 备份时间。
	EndTime string `json:"endTime"`

	// 备份文件大小，单位为MB。
	Size int64 `json:"size"`

	// 备份时间，单位为秒。
	Duration int64 `json:"duration"`

	// 是否加密。true表示加密，默认值为\"false\"，不加密。
	Encrypted *bool `json:"encrypted,omitempty"`
}

func (o Backup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Backup struct{}"
	}

	return strings.Join([]string{"Backup", string(data)}, " ")
}
