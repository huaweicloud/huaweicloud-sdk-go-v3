package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PolicyConfig 策略信息体。设置迁移、同步策略，包括冲突策略、过滤DROP Datase、对象同步范围等。
type PolicyConfig struct {

	// **参数解释**：  过滤DDL策略，指定增量同步阶段需要过滤的DDL操作类型。  **约束限制**：  - 实时迁移场景：MySQL迁移可填空，表示不过滤DROP DATABASE。 - 实时同步场景：MySQL同步只能填drop_database。  **取值范围**：  - drop_database：过滤DROP DATABASE操作  **默认取值**：  不涉及
	FilterDdlPolicy *PolicyConfigFilterDdlPolicy `json:"filter_ddl_policy,omitempty"`

	// **参数解释**：  增量阶段冲突策略，指定当目标库已存在相同数据时的处理方式。该冲突策略特指增量同步中的冲突处理策略，全量阶段的冲突默认忽略。  **约束限制**：  - 实时迁移场景：不支持。 - 实时同步场景：支持。  **取值范围**：  - ignore：忽略，保留目标库中的冲突数据，继续进行后续同步 - stop：报错，同步任务将失败并立即中止 - overwrite：覆盖，用源库数据覆盖目标库已有数据  **默认取值**：  不涉及
	ConflictPolicy *PolicyConfigConflictPolicy `json:"conflict_policy,omitempty"`

	// 对象同步范围：是否同步普通索引。DRS将默认同步主键/唯一索引，普通索引是指除主键/唯一索引以外的其他类型索引。取值： - true：将会同步全部的索引。 - false：仅同步主键/唯一索引，普通索引不会同步。
	IndexTrans *bool `json:"index_trans,omitempty"`

	// 对象同步范围：同步增量阶段是否同步DDL。取值： - true：增量阶段同步DDL。 - false：增量阶段不同步DDL。
	DdlTrans *bool `json:"ddl_trans,omitempty"`

	// 数据同步拓扑。数据同步功能支持多种同步拓扑，您可以根据业务需求规划您的同步实例。参考链接。取值： - one2one：一对一。 - one2many：一对多。 - many2one：多对一。
	DataSyncTopologyType *PolicyConfigDataSyncTopologyType `json:"data_sync_topology_type,omitempty"`

	// **参数解释**：  增量支持的DDL，指定增量同步阶段允许同步的DDL操作类型。  **约束限制**：  不涉及。  **取值范围**：  - CREATE_TABLE：创建表 - ADD_COLUMN：加列 - MODIFY_COLUMN：改列属性 - CHANGE_COLUMN：改列属性 - DROP_INDEX：删除索引 - ADD_INDEX：加索引 - CREATE_INDEX：创建索引 - RENAME_INDEX：重命名索引  **默认取值**：  不涉及
	SupportDdlInfo *PolicyConfigSupportDdlInfo `json:"support_ddl_info,omitempty"`

	// **参数解释**：  同步数据类型，指定增量同步阶段需要同步的数据类型范围。  **约束限制**：  不涉及。  **取值范围**：  - supportAllType：同步所有类型 - tableData：同步数据 - tableStructure：同步表结构 - constraintData：索引同步  **默认取值**：  不涉及
	SyncTypePolicy *string `json:"sync_type_policy,omitempty"`

	// **参数解释**：  Oracle-GaussDB增量读取方式，指定Oracle同步到GaussDB时增量数据的读取模式。  **约束限制**：  不涉及。  **取值范围**：  - logminer：Logminer方式 - xstream：Xstream方式  **默认取值**：  不涉及
	IncrementReadMode *string `json:"increment_read_mode,omitempty"`

	// **参数解释**：  DML同步类型，指定增量同步阶段需要同步的数据操作类型，包括插入、更新和删除操作。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	DmlTypes *string `json:"dml_types,omitempty"`

	// **参数解释**：  索引与表结构是否同时建立，指定在同步过程中创建表时是否同时创建索引。  **约束限制**：  不涉及。  **取值范围**：  - true：索引与表结构同时建立 - false：索引与表结构不同时建立  **默认取值**：  不涉及
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
