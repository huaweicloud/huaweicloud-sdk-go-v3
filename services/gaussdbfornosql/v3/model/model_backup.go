package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Backup 备份详情列表。
type Backup struct {

	// 备份ID。
	Id string `json:"id"`

	// 备份描述信息。
	Description string `json:"description"`

	// 备份所属的实例ID。
	InstanceId string `json:"instance_id"`

	// 备份所属的实例名称。
	InstanceName string `json:"instance_name"`

	Datastore *Datastore `json:"datastore"`

	// 备份名称。
	Name string `json:"name"`

	// - 取值为“Auto”，表示自动全量备份。 - 取值为“Manual”，表示手动全量备份。
	Type string `json:"type"`

	// 备份大小，单位：KB。
	Size float64 `json:"size"`

	// 备份状态。取值： - BUILDING：备份中。 - COMPLETED：备份完成。 - FAILED：备份失败。
	Status string `json:"status"`

	// 备份开始时间，为yyyy-mm-ddThh:mm:ssZ字符串格式，T指某个时间的开始，Z指时区偏移量。
	BeginTime string `json:"begin_time"`

	// 备份结束时间，为yyyy-mm-ddThh:mm:ssZ字符串格式，T指某个时间的开始，Z指时区偏移量。
	EndTime string `json:"end_time"`

	// 备份里的库表信息 - 实例级查询时该字段为空，可忽略 - 库表级查询时该字段非空（存在库表级备份的话）
	DatabaseTables *[]QueryDatabaseTableInfo `json:"database_tables,omitempty"`
}

func (o Backup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Backup struct{}"
	}

	return strings.Join([]string{"Backup", string(data)}, " ")
}
