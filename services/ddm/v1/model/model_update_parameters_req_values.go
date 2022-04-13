package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 需要修改的DDM实例参数的集合。
type UpdateParametersReqValues struct {
	// 用于描述多个拆分表的内在数据关联性，用于告知优化器在处理join时，把join下推到MySQL层执行。格式为:[{tb.col1,tb2.col2},{tb.col2,tb3.col1},...]。

	BindTable *string `json:"bind_table,omitempty"`
	// DDM服务端字符集，如果需要存储emoji表情符号，请选择utf8mb4,并设置RDS字符集也为utf8mb4。修改DDM服务端字符集时，DDM服务端字符序必须同步修改为对应类型的值。

	CharacterSetServer *UpdateParametersReqValuesCharacterSetServer `json:"character_set_server,omitempty"`
	// DDM服务端字符序。修改DDM服务端字符序时，DDM服务端字符集必须同步修改为对应类型的值。

	CollationServer *UpdateParametersReqValuesCollationServer `json:"collation_server,omitempty"`
	// 逻辑表扫描时的分片并发执行级别: DATA_NODE: 分库间并行扫描，同一分库内各分片串行扫描;RDS_INSTANCE: RDS实例间并行扫描，同一RDS实例内各分片串行扫描;PHY_TABLE: 各物理分片全部并行扫描。

	ConcurrentExecutionLevel *UpdateParametersReqValuesConcurrentExecutionLevel `json:"concurrent_execution_level,omitempty"`
	// 服务器关闭连接之前等待连接活动的秒数，以秒为单位，取值范围为60-28800，默认值28800，表示服务器关闭连接之前等待28800秒后，关闭连接。

	ConnectionIdleTimeout *string `json:"connection_idle_timeout,omitempty"`
	// 是否开启表回收站。

	EnableTableRecycle *UpdateParametersReqValuesEnableTableRecycle `json:"enable_table_recycle,omitempty"`
	// insert 常量值使用load data执行。

	InsertToLoadData *UpdateParametersReqValuesInsertToLoadData `json:"insert_to_load_data,omitempty"`
	// 在途事务等待时间窗口，以秒为单位，取值范围为0-100，默认值为1，表示服务器关闭前端连接之前等待1秒后关闭连接。

	LiveTransactionTimeoutOnShutdown *string `json:"live_transaction_timeout_on_shutdown,omitempty"`
	// 记录慢查询的最小秒数,以秒为单位，取值范围为0.01-10，默认值为1，表示如果sql执行大于等于1秒就定义为慢sql。

	LongQueryTime *string `json:"long_query_time,omitempty"`
	// 包或任何生成的中间字符串的最大值。包缓冲区初始化为net_buffer_length字节，但需要时可以增长到max_allowed_packet字节。该值默认很小，以捕获大的（可能是错误的）数据包。该值必须设置为1024的倍数。取值范围为1024~1073741824，默认值为16777216。

	MaxAllowedPacket *string `json:"max_allowed_packet,omitempty"`
	// 允许每个DDM节点同时连接RDS的最大客户端总数。0为默认值标识符,实际值等于(RDS的最大连接数-20)/DDM节点数。取值范围为0-10000000。

	MaxBackendConnections *string `json:"max_backend_connections,omitempty"`
	// 允许同时连接的客户端总数。与后端RDS规格及数量有关。以个数为单位，取值范围为10-40000，默认值为20000，表示允许同时连接的客户端总数不能超过40000。

	MaxConnections *string `json:"max_connections,omitempty"`
	// 允许每个DDM节点同时连接RDS的最小客户端总数。默认值为10。取值范围为0-10000000。

	MinBackendConnections *string `json:"min_backend_connections,omitempty"`
	// 是否强制下推查询语句中不含from的语句。

	NotFromPushdown *UpdateParametersReqValuesNotFromPushdown `json:"not_from_pushdown,omitempty"`
	// 主从rds节点延迟时间阈值，以秒为单位，取值范围为0-7200，默认值为30，表示主rds与从rds之间的数据同步时间值不能超过30秒，如果超过30s，读数据指令就不走当前读节点。

	SecondsBehindMaster *string `json:"seconds_behind_master,omitempty"`
	// 开启或关闭SQL审计。

	SqlAudit *UpdateParametersReqValuesSqlAudit `json:"sql_audit,omitempty"`
	// SQL执行超时秒数，以秒为单位，取值范围为100-28800，默认值28800，表示sql执行大于等于28800秒超时。

	SqlExecuteTimeout *string `json:"sql_execute_timeout,omitempty"`
	// DDL语句添加binlog hint。

	SupportDdlBinlogHint *UpdateParametersReqValuesSupportDdlBinlogHint `json:"support_ddl_binlog_hint,omitempty"`
	// XA：XA 事务，保证原子性，保证可见性；FREE：允许多写，不保证原子性，无性能损耗；NO_DTX：单分片事务。

	TransactionPolicy *UpdateParametersReqValuesTransactionPolicy `json:"transaction_policy,omitempty"`
	// 开启或关闭优化器中的极致下推优化功能。

	UltimateOptimize *UpdateParametersReqValuesUltimateOptimize `json:"ultimate_optimize,omitempty"`
}

func (o UpdateParametersReqValues) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateParametersReqValues struct{}"
	}

	return strings.Join([]string{"UpdateParametersReqValues", string(data)}, " ")
}

type UpdateParametersReqValuesCharacterSetServer struct {
	value string
}

type UpdateParametersReqValuesCharacterSetServerEnum struct {
	GBK     UpdateParametersReqValuesCharacterSetServer
	UTF8    UpdateParametersReqValuesCharacterSetServer
	UTF8MB4 UpdateParametersReqValuesCharacterSetServer
}

func GetUpdateParametersReqValuesCharacterSetServerEnum() UpdateParametersReqValuesCharacterSetServerEnum {
	return UpdateParametersReqValuesCharacterSetServerEnum{
		GBK: UpdateParametersReqValuesCharacterSetServer{
			value: "gbk",
		},
		UTF8: UpdateParametersReqValuesCharacterSetServer{
			value: "utf8",
		},
		UTF8MB4: UpdateParametersReqValuesCharacterSetServer{
			value: "utf8mb4",
		},
	}
}

func (c UpdateParametersReqValuesCharacterSetServer) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateParametersReqValuesCharacterSetServer) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type UpdateParametersReqValuesCollationServer struct {
	value string
}

type UpdateParametersReqValuesCollationServerEnum struct {
	UTF8_UNICODE_CI    UpdateParametersReqValuesCollationServer
	UTF8_BIN           UpdateParametersReqValuesCollationServer
	GBK_CHINESE_CI     UpdateParametersReqValuesCollationServer
	GBK_BIN            UpdateParametersReqValuesCollationServer
	UTF8MB4_UNICODE_CI UpdateParametersReqValuesCollationServer
	UTF8MB4_BIN        UpdateParametersReqValuesCollationServer
}

func GetUpdateParametersReqValuesCollationServerEnum() UpdateParametersReqValuesCollationServerEnum {
	return UpdateParametersReqValuesCollationServerEnum{
		UTF8_UNICODE_CI: UpdateParametersReqValuesCollationServer{
			value: "utf8_unicode_ci",
		},
		UTF8_BIN: UpdateParametersReqValuesCollationServer{
			value: "utf8_bin",
		},
		GBK_CHINESE_CI: UpdateParametersReqValuesCollationServer{
			value: "gbk_chinese_ci",
		},
		GBK_BIN: UpdateParametersReqValuesCollationServer{
			value: "gbk_bin",
		},
		UTF8MB4_UNICODE_CI: UpdateParametersReqValuesCollationServer{
			value: "utf8mb4_unicode_ci",
		},
		UTF8MB4_BIN: UpdateParametersReqValuesCollationServer{
			value: "utf8mb4_bin",
		},
	}
}

func (c UpdateParametersReqValuesCollationServer) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateParametersReqValuesCollationServer) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type UpdateParametersReqValuesConcurrentExecutionLevel struct {
	value string
}

type UpdateParametersReqValuesConcurrentExecutionLevelEnum struct {
	RDS_INSTANCE UpdateParametersReqValuesConcurrentExecutionLevel
	DATA_NODE    UpdateParametersReqValuesConcurrentExecutionLevel
	PHY_TABLE    UpdateParametersReqValuesConcurrentExecutionLevel
}

func GetUpdateParametersReqValuesConcurrentExecutionLevelEnum() UpdateParametersReqValuesConcurrentExecutionLevelEnum {
	return UpdateParametersReqValuesConcurrentExecutionLevelEnum{
		RDS_INSTANCE: UpdateParametersReqValuesConcurrentExecutionLevel{
			value: "RDS_INSTANCE",
		},
		DATA_NODE: UpdateParametersReqValuesConcurrentExecutionLevel{
			value: "DATA_NODE",
		},
		PHY_TABLE: UpdateParametersReqValuesConcurrentExecutionLevel{
			value: "PHY_TABLE",
		},
	}
}

func (c UpdateParametersReqValuesConcurrentExecutionLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateParametersReqValuesConcurrentExecutionLevel) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type UpdateParametersReqValuesEnableTableRecycle struct {
	value string
}

type UpdateParametersReqValuesEnableTableRecycleEnum struct {
	OFF UpdateParametersReqValuesEnableTableRecycle
	ON  UpdateParametersReqValuesEnableTableRecycle
}

func GetUpdateParametersReqValuesEnableTableRecycleEnum() UpdateParametersReqValuesEnableTableRecycleEnum {
	return UpdateParametersReqValuesEnableTableRecycleEnum{
		OFF: UpdateParametersReqValuesEnableTableRecycle{
			value: "OFF",
		},
		ON: UpdateParametersReqValuesEnableTableRecycle{
			value: "ON",
		},
	}
}

func (c UpdateParametersReqValuesEnableTableRecycle) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateParametersReqValuesEnableTableRecycle) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type UpdateParametersReqValuesInsertToLoadData struct {
	value string
}

type UpdateParametersReqValuesInsertToLoadDataEnum struct {
	OFF UpdateParametersReqValuesInsertToLoadData
	ON  UpdateParametersReqValuesInsertToLoadData
}

func GetUpdateParametersReqValuesInsertToLoadDataEnum() UpdateParametersReqValuesInsertToLoadDataEnum {
	return UpdateParametersReqValuesInsertToLoadDataEnum{
		OFF: UpdateParametersReqValuesInsertToLoadData{
			value: "OFF",
		},
		ON: UpdateParametersReqValuesInsertToLoadData{
			value: "ON",
		},
	}
}

func (c UpdateParametersReqValuesInsertToLoadData) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateParametersReqValuesInsertToLoadData) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type UpdateParametersReqValuesNotFromPushdown struct {
	value string
}

type UpdateParametersReqValuesNotFromPushdownEnum struct {
	OFF UpdateParametersReqValuesNotFromPushdown
	ON  UpdateParametersReqValuesNotFromPushdown
}

func GetUpdateParametersReqValuesNotFromPushdownEnum() UpdateParametersReqValuesNotFromPushdownEnum {
	return UpdateParametersReqValuesNotFromPushdownEnum{
		OFF: UpdateParametersReqValuesNotFromPushdown{
			value: "OFF",
		},
		ON: UpdateParametersReqValuesNotFromPushdown{
			value: "ON",
		},
	}
}

func (c UpdateParametersReqValuesNotFromPushdown) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateParametersReqValuesNotFromPushdown) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type UpdateParametersReqValuesSqlAudit struct {
	value string
}

type UpdateParametersReqValuesSqlAuditEnum struct {
	OFF UpdateParametersReqValuesSqlAudit
	ON  UpdateParametersReqValuesSqlAudit
}

func GetUpdateParametersReqValuesSqlAuditEnum() UpdateParametersReqValuesSqlAuditEnum {
	return UpdateParametersReqValuesSqlAuditEnum{
		OFF: UpdateParametersReqValuesSqlAudit{
			value: "OFF",
		},
		ON: UpdateParametersReqValuesSqlAudit{
			value: "ON",
		},
	}
}

func (c UpdateParametersReqValuesSqlAudit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateParametersReqValuesSqlAudit) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type UpdateParametersReqValuesSupportDdlBinlogHint struct {
	value string
}

type UpdateParametersReqValuesSupportDdlBinlogHintEnum struct {
	OFF UpdateParametersReqValuesSupportDdlBinlogHint
	ON  UpdateParametersReqValuesSupportDdlBinlogHint
}

func GetUpdateParametersReqValuesSupportDdlBinlogHintEnum() UpdateParametersReqValuesSupportDdlBinlogHintEnum {
	return UpdateParametersReqValuesSupportDdlBinlogHintEnum{
		OFF: UpdateParametersReqValuesSupportDdlBinlogHint{
			value: "OFF",
		},
		ON: UpdateParametersReqValuesSupportDdlBinlogHint{
			value: "ON",
		},
	}
}

func (c UpdateParametersReqValuesSupportDdlBinlogHint) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateParametersReqValuesSupportDdlBinlogHint) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type UpdateParametersReqValuesTransactionPolicy struct {
	value string
}

type UpdateParametersReqValuesTransactionPolicyEnum struct {
	XA     UpdateParametersReqValuesTransactionPolicy
	FREE   UpdateParametersReqValuesTransactionPolicy
	NO_DTX UpdateParametersReqValuesTransactionPolicy
}

func GetUpdateParametersReqValuesTransactionPolicyEnum() UpdateParametersReqValuesTransactionPolicyEnum {
	return UpdateParametersReqValuesTransactionPolicyEnum{
		XA: UpdateParametersReqValuesTransactionPolicy{
			value: "XA",
		},
		FREE: UpdateParametersReqValuesTransactionPolicy{
			value: "FREE",
		},
		NO_DTX: UpdateParametersReqValuesTransactionPolicy{
			value: "NO_DTX",
		},
	}
}

func (c UpdateParametersReqValuesTransactionPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateParametersReqValuesTransactionPolicy) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type UpdateParametersReqValuesUltimateOptimize struct {
	value string
}

type UpdateParametersReqValuesUltimateOptimizeEnum struct {
	OFF UpdateParametersReqValuesUltimateOptimize
	ON  UpdateParametersReqValuesUltimateOptimize
}

func GetUpdateParametersReqValuesUltimateOptimizeEnum() UpdateParametersReqValuesUltimateOptimizeEnum {
	return UpdateParametersReqValuesUltimateOptimizeEnum{
		OFF: UpdateParametersReqValuesUltimateOptimize{
			value: "OFF",
		},
		ON: UpdateParametersReqValuesUltimateOptimize{
			value: "ON",
		},
	}
}

func (c UpdateParametersReqValuesUltimateOptimize) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateParametersReqValuesUltimateOptimize) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
