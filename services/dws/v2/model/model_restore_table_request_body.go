package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreTableRequestBody **参数解释**： 恢复表请求体。 **取值范围**： 不涉及。
type RestoreTableRequestBody struct {

	// **参数解释**： 名称是否区分大小写。 **取值范围**： 不涉及。
	CaseSensitive bool `json:"case_sensitive"`

	// **参数解释**： 数据库名称。 **取值范围**： 不涉及。
	Database string `json:"database"`

	// **参数解释**： 源表信息。 **取值范围**： 不涉及。
	RestoreTableList []TableDetail `json:"restore_table_list"`

	// **参数解释**： 目的表信息。 **取值范围**： 不涉及。
	TargetTableList []TableDetail `json:"target_table_list"`
}

func (o RestoreTableRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreTableRequestBody struct{}"
	}

	return strings.Join([]string{"RestoreTableRequestBody", string(data)}, " ")
}
