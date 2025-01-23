package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRestoreTablesRequest Request Object
type ShowRestoreTablesRequest struct {

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// 备份时间点，时间戳格式。  通过[查询可恢复时间段](https://support.huaweicloud.com/api-taurusdb/ShowBackupRestoreTime.html)获取。
	RestoreTime string `json:"restore_time"`

	// 是否是最新库表。 - true：是最新库表。 - false：是恢复时间点库表。
	LastTableInfo string `json:"last_table_info"`

	// 数据库名称，模糊匹配。
	DatabaseName *string `json:"database_name,omitempty"`

	// 表名称，模糊匹配。
	TableName *string `json:"table_name,omitempty"`
}

func (o ShowRestoreTablesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRestoreTablesRequest struct{}"
	}

	return strings.Join([]string{"ShowRestoreTablesRequest", string(data)}, " ")
}
