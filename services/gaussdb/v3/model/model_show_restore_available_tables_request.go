package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRestoreAvailableTablesRequest Request Object
type ShowRestoreAvailableTablesRequest struct {

	// **参数解释**：              请求语言类型。  **约束限制**：  不涉及。  **取值范围**： - en-us。 - zh-cn。  **默认值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in07，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：              备份时间点。  **约束限制**：  格式为UNIX时间戳，单位是毫秒，时区为UTC标准时区。传参时需要将对应时区的时间转为标准时区对应的时间戳，比如，北京时区的时间点需要-8h后再转为时间戳。  **取值范围**：              通过[查询可恢复时间段](https://support.huaweicloud.com/api-taurusdb/ShowBackupRestoreTime.html)获取。  **默认取值**：  不涉及。
	RestoreTime string `json:"restore_time"`

	// **参数解释**：              是否是最新库表。  **约束限制**：  不涉及。  **取值范围**：  - true：是最新库表。 - false：是恢复时间点库表。  **默认取值**：  不涉及。
	LastTableInfo string `json:"last_table_info"`

	// **参数解释**：   数据库名称，模糊匹配。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	DatabaseName *string `json:"database_name,omitempty"`

	// **参数解释**：   表名称，模糊匹配。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	TableName *string `json:"table_name,omitempty"`
}

func (o ShowRestoreAvailableTablesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRestoreAvailableTablesRequest struct{}"
	}

	return strings.Join([]string{"ShowRestoreAvailableTablesRequest", string(data)}, " ")
}
