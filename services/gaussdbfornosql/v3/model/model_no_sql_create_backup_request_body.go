package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NoSqlCreateBackupRequestBody **参数解释：** 创建手动备份请求参数。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
type NoSqlCreateBackupRequestBody struct {

	// **参数解释：** 手动备份名称。 **约束限制：** 长度为4~64位，必须以字母开头（A~Z或a~z），区分大小写，可以包含字母、数字（0~9）、中划线（-）或者下划线（_），不能包含其他特殊字符。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Name string `json:"name"`

	// **参数解释：** 手动备份描述。 **约束限制：** 长度不超过256位，且不能包含>!<\"&'=特殊字符。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Description string `json:"description"`

	// **参数解释：** 备份里的库表信息。 **约束限制：** 该参数仅针对GeminiDB Cassandra。 **取值范围：** - 字段为空，表示创建实例级备份。 - 字段非空，表示创建库表级备份。  **默认取值：** 不涉及。
	DatabaseTables *[]DatabaseTable `json:"database_tables,omitempty"`
}

func (o NoSqlCreateBackupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NoSqlCreateBackupRequestBody struct{}"
	}

	return strings.Join([]string{"NoSqlCreateBackupRequestBody", string(data)}, " ")
}
