package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PolicyConfig 策略信息体。设置迁移、同步策略，包括冲突策略、过滤DROP Datase、对象同步范围等。
type PolicyConfig struct {

	// 过滤DDL策略。取值： - drop_database  场景区别： - 实时迁移场景：MySQL迁移可填\"\"，表示不过滤DROP DATABASE。 - 实时同步场景：MySQL同步只能填\"drop_database\"。
	FilterDdlPolicy *PolicyConfigFilterDdlPolicy `json:"filter_ddl_policy,omitempty"`

	// 增量阶段冲突策略。该冲突策略特指增量同步中的冲突处理策略,全量阶段的冲突默认忽略。取值： - ignore：忽略。当同步数据与目标库已有数据冲突时（主键/唯一键重复等），DRS将忽略源库的冲突数据，并保留目标库中的冲突数据，继续进行后续同步。  - stop：报错。当同步数据与目标库已有数据冲突时（主键/唯一键重复等），同步任务将失败并立即中止。可在同步日志中查看详细信息。  - overwrite：覆盖。当同步数据与目标库已有数据冲突时（主键/唯一键重复等），将覆盖原来的冲突数据。  场景区别： - 实时迁移场景：不支持。 - 实时同步场景：支持。
	ConflictPolicy *PolicyConfigConflictPolicy `json:"conflict_policy,omitempty"`

	// 对象同步范围：是否同步普通索引。DRS将默认同步主键/唯一索引，普通索引是指除主键/唯一索引以外的其他类型索引。取值： - true：将会同步全部的索引。 - false：仅同步主键/唯一索引，普通索引不会同步。
	IndexTrans *bool `json:"index_trans,omitempty"`

	// 对象同步范围：同步增量阶段是否同步DDL。取值： - true：增量阶段同步DDL。 - false：增量阶段不同步DDL。
	DdlTrans *bool `json:"ddl_trans,omitempty"`

	// 数据同步拓扑。数据同步功能支持多种同步拓扑，您可以根据业务需求规划您的同步实例。参考链接。取值： - one2one：一对一。 - one2many：一对多。 - many2one：多对一。
	DataSyncTopologyType *PolicyConfigDataSyncTopologyType `json:"data_sync_topology_type,omitempty"`

	// 增量支持的DDL。取值： - \"CREATE_TABLE,ADD_COLUMN,MODIFY_COLUMN,CHANGE_COLUMN,DROP_INDEX,ADD_INDEX,CREATE_INDEX,RENAME_INDEX\"。 - 含义解释： - CREATE_TABLE：创建表。 - ADD_COLUMN：加列。 - MODIFY_COLUMN：改列属性。 - CHANGE_COLUMN：改列属性。 - DROP_INDEX：删除索引。 - ADD_INDEX：加索引。 - CREATE_INDEX：创建索引。 - RENAME_INDEX：重命名索引。 - 使用提示： 1.一对一，一对多场景，如果业务上认为源和目标应该使用保持严格一致，那么高危类DDL也应该勾选并同步。 2.一对一，一对多场景，如果业务上确定某个高危DDL不应该发生，则可以不勾选同步高危类DDL，这样DRS将拦截过滤这个DDL，从而起到保护目标数据的作用，但需要知晓过滤DDL的附带问题是可能导致同步失败，例如过滤删列动作。 3.多对一数据聚合场景，最佳实践是推荐只选择同步加列DDL，其他大部分DDL同步都因目标表修改而导致其他任务失败/数据不一致的情况发生，常见情况有：a、同步truncate导致目标数据全部被清空； b、同步创建索引导致目标表被锁定； c、同步rename导致其他任务找不到目标表而失败；d、同步改列导致其他任务因数据类型不兼容而失败；
	SupportDdlInfo *PolicyConfigSupportDdlInfo `json:"support_ddl_info,omitempty"`

	// 同步数据类型。取值：supportAllType（同步所有类型），tableData（同步数据），tableStructure（同步表结构），constraintData（索引同步）。 说明：除supportAllType以外，其他类型可组合填写，例如：\"tableData,tableStructure\" 。
	SyncTypePolicy *string `json:"sync_type_policy,omitempty"`

	// oracle-gausssdb增量读取方式：logminer，xstream
	IncrementReadMode *string `json:"increment_read_mode,omitempty"`

	// DML同步类型。
	DmlTypes *string `json:"dml_types,omitempty"`

	// 索引与表结构是否同时建立。
	IsCreateTableWithIndex *bool `json:"is_create_table_with_index,omitempty"`
}

func (o PolicyConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyConfig struct{}"
	}

	return strings.Join([]string{"PolicyConfig", string(data)}, " ")
}

type PolicyConfigFilterDdlPolicy struct {
	value string
}

type PolicyConfigFilterDdlPolicyEnum struct {
	DROP_DATABASE PolicyConfigFilterDdlPolicy
}

func GetPolicyConfigFilterDdlPolicyEnum() PolicyConfigFilterDdlPolicyEnum {
	return PolicyConfigFilterDdlPolicyEnum{
		DROP_DATABASE: PolicyConfigFilterDdlPolicy{
			value: "drop_database",
		},
	}
}

func (c PolicyConfigFilterDdlPolicy) Value() string {
	return c.value
}

func (c PolicyConfigFilterDdlPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyConfigFilterDdlPolicy) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type PolicyConfigConflictPolicy struct {
	value string
}

type PolicyConfigConflictPolicyEnum struct {
	IGNORE    PolicyConfigConflictPolicy
	STOP      PolicyConfigConflictPolicy
	OVERWRITE PolicyConfigConflictPolicy
}

func GetPolicyConfigConflictPolicyEnum() PolicyConfigConflictPolicyEnum {
	return PolicyConfigConflictPolicyEnum{
		IGNORE: PolicyConfigConflictPolicy{
			value: "ignore",
		},
		STOP: PolicyConfigConflictPolicy{
			value: "stop",
		},
		OVERWRITE: PolicyConfigConflictPolicy{
			value: "overwrite",
		},
	}
}

func (c PolicyConfigConflictPolicy) Value() string {
	return c.value
}

func (c PolicyConfigConflictPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyConfigConflictPolicy) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type PolicyConfigDataSyncTopologyType struct {
	value string
}

type PolicyConfigDataSyncTopologyTypeEnum struct {
	ONE2ONE  PolicyConfigDataSyncTopologyType
	ONE2MANY PolicyConfigDataSyncTopologyType
	MANY2ONE PolicyConfigDataSyncTopologyType
}

func GetPolicyConfigDataSyncTopologyTypeEnum() PolicyConfigDataSyncTopologyTypeEnum {
	return PolicyConfigDataSyncTopologyTypeEnum{
		ONE2ONE: PolicyConfigDataSyncTopologyType{
			value: "one2one",
		},
		ONE2MANY: PolicyConfigDataSyncTopologyType{
			value: "one2many",
		},
		MANY2ONE: PolicyConfigDataSyncTopologyType{
			value: "many2one",
		},
	}
}

func (c PolicyConfigDataSyncTopologyType) Value() string {
	return c.value
}

func (c PolicyConfigDataSyncTopologyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyConfigDataSyncTopologyType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type PolicyConfigSupportDdlInfo struct {
	value string
}

type PolicyConfigSupportDdlInfoEnum struct {
	CREATE_TABLE  PolicyConfigSupportDdlInfo
	ADD_COLUMN    PolicyConfigSupportDdlInfo
	MODIFY_COLUMN PolicyConfigSupportDdlInfo
	CHANGE_COLUMN PolicyConfigSupportDdlInfo
	DROP_INDEX    PolicyConfigSupportDdlInfo
	ADD_INDEX     PolicyConfigSupportDdlInfo
	CREATE_INDEX  PolicyConfigSupportDdlInfo
	RENAME_INDEX  PolicyConfigSupportDdlInfo
}

func GetPolicyConfigSupportDdlInfoEnum() PolicyConfigSupportDdlInfoEnum {
	return PolicyConfigSupportDdlInfoEnum{
		CREATE_TABLE: PolicyConfigSupportDdlInfo{
			value: "CREATE_TABLE",
		},
		ADD_COLUMN: PolicyConfigSupportDdlInfo{
			value: "ADD_COLUMN",
		},
		MODIFY_COLUMN: PolicyConfigSupportDdlInfo{
			value: "MODIFY_COLUMN",
		},
		CHANGE_COLUMN: PolicyConfigSupportDdlInfo{
			value: "CHANGE_COLUMN",
		},
		DROP_INDEX: PolicyConfigSupportDdlInfo{
			value: "DROP_INDEX",
		},
		ADD_INDEX: PolicyConfigSupportDdlInfo{
			value: "ADD_INDEX",
		},
		CREATE_INDEX: PolicyConfigSupportDdlInfo{
			value: "CREATE_INDEX",
		},
		RENAME_INDEX: PolicyConfigSupportDdlInfo{
			value: "RENAME_INDEX",
		},
	}
}

func (c PolicyConfigSupportDdlInfo) Value() string {
	return c.value
}

func (c PolicyConfigSupportDdlInfo) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyConfigSupportDdlInfo) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
