package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DdlInfoItem struct {

	// **参数解释**：   无锁变更产生的临时表，获取方法参见[查询无锁变更任务记录列表](https://support.huaweicloud.com/api-taurusdb/ListOnlineDdlTaskRecords.html)的响应参数temp_table_name。   **约束限制**：  不涉及。   **取值范围**：   不涉及。   **默认取值**：   不涉及。
	Table string `json:"table"`
}

func (o DdlInfoItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DdlInfoItem struct{}"
	}

	return strings.Join([]string{"DdlInfoItem", string(data)}, " ")
}
