package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskContentItem struct {

	// **参数解释**：  无锁变更的目标数据库。  **约束限制**： 不涉及。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	Schema string `json:"schema"`

	// **参数解释**：  无锁变更的DDL信息，包含临时表名。  **约束限制**： 不涉及。
	DdlInfo []DdlInfoItem `json:"ddl_info"`
}

func (o TaskContentItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskContentItem struct{}"
	}

	return strings.Join([]string{"TaskContentItem", string(data)}, " ")
}
