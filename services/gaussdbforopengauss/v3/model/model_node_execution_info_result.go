package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeExecutionInfoResult SQL链路节点执行信息
type NodeExecutionInfoResult struct {

	// **参数解释**: 组件ID。 **取值范围**: 不涉及。
	ComponentId string `json:"component_id"`

	// **参数解释**: 节点ID。 **取值范围**: 不涉及。
	NodeId string `json:"node_id"`

	// **参数解释**: 事务ID。 **取值范围**: 不涉及。
	TransactionId string `json:"transaction_id"`

	// **参数解释**: 归一化SQL ID。 **取值范围**: 不涉及。
	SqlId string `json:"sql_id"`

	// **参数解释**: 唯一SQL ID。 **取值范围**: 不涉及。
	SqlExecId string `json:"sql_exec_id"`

	// **参数解释**: 数据库名称。 **取值范围**: 不涉及。
	DbName string `json:"db_name"`

	// **参数解释**: schema名称。 **取值范围**: 不涉及。
	SchemaName string `json:"schema_name"`

	// **参数解释**: 语句启动的时间，格式为“yyyy-mm-ddThh:mm: ssssssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。 **取值范围**: 不涉及。
	StartTime string `json:"start_time"`

	// **参数解释**: 语句结束的时间，格式为“yyyy-mm-ddThh:mm: ssssssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。 **取值范围**: 不涉及。
	FinishTime string `json:"finish_time"`

	// **参数解释**: 执行总耗时（单位：微秒）。 **取值范围**: 不涉及。
	AllTime int64 `json:"all_time"`

	// **参数解释**: 用户名。 **取值范围**: 不涉及。
	UserName string `json:"user_name"`

	// **参数解释**: 用户发起的请求的客户端地址。 **取值范围**: 不涉及。
	ClientAddr string `json:"client_addr"`

	// **参数解释**: 用户发起的请求的客户端端口。 **取值范围**: 不涉及。
	ClientPort int32 `json:"client_port"`

	// **参数解释**: 驱动传入的trace id，与应用的一次请求相关联。 **取值范围**: 不涉及。
	TraceId string `json:"trace_id"`

	// **参数解释**: 用户发起的请求的应用程序名称。 **取值范围**: 不涉及。
	ApplicationName string `json:"application_name"`

	// **参数解释**: 用户session id。 **取值范围**: 不涉及。
	SessionId string `json:"session_id"`

	// **参数解释**: 该SQL是否为slow SQL。 **取值范围**: 不涉及。
	IsSlowSql *bool `json:"is_slow_sql,omitempty"`

	ExecutionTimeDetails *ExecutionTimeDetailsInfo `json:"execution_time_details"`
}

func (o NodeExecutionInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeExecutionInfoResult struct{}"
	}

	return strings.Join([]string{"NodeExecutionInfoResult", string(data)}, " ")
}
