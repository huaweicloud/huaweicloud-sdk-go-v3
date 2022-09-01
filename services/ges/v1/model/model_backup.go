package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 图备份
type Backup struct {

	// 备份ID。
	Id string `json:"id" xml:"id"`

	// 备份名称。
	Name string `json:"name" xml:"name"`

	// 备份方法，取值为auto或manual。
	BackupMethod string `json:"backupMethod" xml:"backupMethod"`

	// 备份关联的图ID。
	GraphId string `json:"graphId" xml:"graphId"`

	// 备份关联的图Name。
	GraphName string `json:"graphName" xml:"graphName"`

	// 备份关联的图状态。
	GraphStatus string `json:"graphStatus" xml:"graphStatus"`

	// 备份关联的图规格。
	GraphSizeTypeIndex string `json:"graphSizeTypeIndex" xml:"graphSizeTypeIndex"`

	// 备份关联的图版本。
	DataStoreVersion string `json:"dataStoreVersion" xml:"dataStoreVersion"`

	// 备份关联的图CPU架构。
	Arch string `json:"arch" xml:"arch"`

	// 备份状态。  - backing_up：备份中 - success：备份成功 - failed：备份失败
	Status string `json:"status" xml:"status"`

	// 备份开始时间戳。
	StartTimestamp int64 `json:"startTimestamp" xml:"startTimestamp"`

	// 备份时间。
	StartTime string `json:"startTime" xml:"startTime"`

	// 备份结束时间戳。
	EndTimestamp int64 `json:"endTimestamp" xml:"endTimestamp"`

	// 备份时间。
	EndTime string `json:"endTime" xml:"endTime"`

	// 备份文件大小，单位为MB。
	Size int64 `json:"size" xml:"size"`

	// 备份时间，单位为秒。
	Duration int64 `json:"duration" xml:"duration"`

	// 是否加密。true表示加密，默认值为\"false\"，不加密。
	Encrypted *bool `json:"encrypted,omitempty" xml:"encrypted"`
}

func (o Backup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Backup struct{}"
	}

	return strings.Join([]string{"Backup", string(data)}, " ")
}
