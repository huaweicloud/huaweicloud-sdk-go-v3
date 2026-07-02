package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RealtimeSessionRequest 收集全部实时会话信息请求体
type RealtimeSessionRequest struct {

	// **参数解释**：  需要收集的实时会话类型。  **约束限制**：  不涉及。  **取值范围**：  - slow：慢会话。 - active：活跃会话。 - total：会话总数。 - long：长事务会话。  **默认取值**：  total。
	Filter *string `json:"filter,omitempty"`

	// **参数解释**：  慢会话阈值，单位是秒。  **约束限制**：  不涉及。  **取值范围**：  1-86400。  **默认取值**：  10。
	SlowProcessThreshold *int32 `json:"slow_process_threshold,omitempty"`

	// **参数解释**：  实时会话的用户。 获取方法请参见[查询数据库用户](https://support.huaweicloud.com/api-taurusdb/ListGaussMySqlDatabaseUser.html)。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	User *string `json:"user,omitempty"`

	// **参数解释**：  实时会话的主机。 获取方法请参见[查询数据库用户](https://support.huaweicloud.com/api-taurusdb/ListGaussMySqlDatabaseUser.html)。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Host *string `json:"host,omitempty"`

	// **参数解释**：  实时会话的数据库。 获取方法请参见[查询数据库用户](https://support.huaweicloud.com/api-taurusdb/ListGaussMySqlDatabaseUser.html)。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Db *string `json:"db,omitempty"`

	// **参数解释**：  实时会话的命令类型。  **约束限制**：  不涉及。  **取值范围**：  - Sleep：空闲连接，无任何操作。 - Query：正在执行查询。 - Connect：建立连接。 - Init DB：切换数据库。 - Field List：获取表字段列表。 - Processlist：查看会话列表。  **默认取值**：  不涉及。
	Command *string `json:"command,omitempty"`

	// **参数解释**：  实时会话的SQL语句。 您可以通过登录管理控制台，选择智能DBA助手下的实时会话，在会话列表中获取。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	SqlKey *string `json:"sql_key,omitempty"`

	// **参数解释**：  实时会话的排序列。  **约束限制**：  不涉及。  **取值范围**：  - id：会话ID。 - state_duration：状态持续时间。 - trx_executed_time：事务持续时间。 - trx_id：事务ID。 - trx_lock_duration：事务锁等待时长。 - trx_lock_rows：事务锁定行数。 - trx_lock_tables：事务锁定表数量。 - trx_update_rows：事务更新行数。  **默认取值**：  id。
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释**：  实时会话的排序方向。  **约束限制**：  不涉及。  **取值范围**：  - desc：降序排列。 - asc：升序排列。  **默认取值**：  asc。
	SortDir *string `json:"sort_dir,omitempty"`
}

func (o RealtimeSessionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RealtimeSessionRequest struct{}"
	}

	return strings.Join([]string{"RealtimeSessionRequest", string(data)}, " ")
}
