package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StartOnlineTaskContentItem struct {

	// **参数解释**：  无锁变更的目标数据库。 获取方法请参见[查询数据库列表](https://support.huaweicloud.com/api-taurusdb/ListGaussMySqlDatabase.html)。  **约束限制**：  不涉及。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	Schema string `json:"schema"`

	// **参数解释**：  无锁变更的DDL信息。  **约束限制**： 不涉及。
	DdlInfo []StartOnlineDdlInfoItem `json:"ddl_info"`
}

func (o StartOnlineTaskContentItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartOnlineTaskContentItem struct{}"
	}

	return strings.Join([]string{"StartOnlineTaskContentItem", string(data)}, " ")
}
