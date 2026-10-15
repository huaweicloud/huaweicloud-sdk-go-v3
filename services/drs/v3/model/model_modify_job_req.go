package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ModifyJobReq 修改任务描述信息、名称，设置异常通知信息。
type ModifyJobReq struct {

	// **参数解释**：  迁移任务ID，用于唯一标识一个数据复制任务，在修改任务时作为任务的识别凭证。  **约束限制**：  不涉及。  **取值范围**：  UUID格式字符串。  **默认取值**：  不涉及
	JobId string `json:"job_id"`

	// **参数解释**：  任务描述，用于补充说明任务的用途或备注信息。  **约束限制**：  - 修改任务描述时必填。 - 长度为0-256个字符。 - 不能包含!<>&'\"特殊字符。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**：  任务名称，用于标识和区分不同的数据复制任务。  **约束限制**：  - 需要修改任务名称时填写。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Name *string `json:"name,omitempty"`

	AlarmNotify *AlarmNotifyInfo `json:"alarm_notify,omitempty"`

	// **参数解释**：  任务迁移模式，指定数据复制任务的执行方式。  **约束限制**：  - 目前仅支持MySQL到TaurusDB同步任务在配置中或者增量中时修改。  **取值范围**：  - FULL_TRANS：全量迁移，仅迁移历史数据 - FULL_INCR_TRANS：全量+增量迁移，先迁移历史数据再持续同步增量 - INCR_TRANS：增量迁移，仅同步数据变更  **默认取值**：  不涉及
	TaskType *ModifyJobReqTaskType `json:"task_type,omitempty"`

	SourceEndpoint *Endpoint `json:"source_endpoint,omitempty"`

	TargetEndpoint *Endpoint `json:"target_endpoint,omitempty"`

	// **参数解释**：  DRS任务实例规格编码，决定迁移实例的计算资源大小，测试连接之后修改调用时必填。  **约束限制**：  - 测试连接之后修改调用时必填。 - 具体某种场景支持的取值可以通过查询可用的Node规格获取。  **取值范围**：  - micro：极小规格 - small：小规格 - medium：中规格 - high：大规格 - xlarge：超大规格 - 2xlarge：极大规格  **默认取值**：  high
	NodeType *ModifyJobReqNodeType `json:"node_type,omitempty"`

	// **参数解释**：  DRS任务引擎类型，测试连接之后修改调用时必填。  **约束限制**：  - 测试连接之后修改调用时必填。 - 详细取值可参考引擎类型说明。  **取值范围**：  - mysql：MySQL到MySQL迁移，MySQL到MySQL同步 - mongodb：MongoDB到DDS迁移与同步 - cloudDataGuard-mysql：MySQL到MySQL灾备 - gaussdbv5：GaussDB同步 - mysql-to-kafka：MySQL到Kafka同步 - taurus-to-kafka：TaurusDB到Kafka同步 - gaussdbv5ha-to-kafka：GaussDB集中式版到Kafka同步 - postgresql：PostgreSQL到PostgreSQL同步 - oracle-to-gaussdbv5：Oracle-GaussDB分布式版 - oracle-to-gaussdbv5ha：Oracle-GaussDB集中式版 - gaussdbv5-to-oracle：GaussDB分布式版-Oracle - gaussdbv5ha-to-oracle：GaussDB集中式版-Oracle - gaussdbv5-to-kafka：GaussDB分布式版-Kafka  **默认取值**：  不涉及
	EngineType *ModifyJobReqEngineType `json:"engine_type,omitempty"`

	// **参数解释**：  任务使用的网络类型，测试连接之后调用必填。  **约束限制**：  - 测试连接之后调用必填。  **取值范围**：  - vpn：虚拟专用网络 - vpc：虚拟私有云 - eip：弹性公网IP  **默认取值**：  不涉及
	NetType *ModifyJobReqNetType `json:"net_type,omitempty"`

	// **参数解释**：  是否保存数据库信息，测试连接之后调用必填。  **约束限制**：  - 测试连接之后调用必填。  **取值范围**：  - true：保存数据库信息 - false：不保存数据库信息  **默认取值**：  不涉及
	StoreDbInfo *bool `json:"store_db_info,omitempty"`

	// **参数解释**：  是否为重建任务。  **约束限制**：  不涉及。  **取值范围**：  - true：重建任务 - false：非重建任务  **默认取值**：  不涉及
	IsRecreate *bool `json:"is_recreate,omitempty"`

	// **参数解释**：  任务的迁移方向，测试连接之后调用必填。  **约束限制**：  - 测试连接之后调用必填。  **取值范围**：  - up：入云，灾备场景时对应本云为备 - down：出云，灾备场景时对应本云为主 - non-dbs：自建  **默认取值**：  不涉及
	JobDirection *ModifyJobReqJobDirection `json:"job_direction,omitempty"`

	// **参数解释**：  指定目标数据库实例在迁移过程中是否限制为只读状态。  **约束限制**：  不涉及。  **取值范围**：  - true：目标实例限制为只读 - false：目标实例不限制为只读  **默认取值**：  不涉及
	IsTargetReadonly *bool `json:"is_target_readonly,omitempty"`

	// **参数解释**：  所有Definer是否迁移到该用户下，MySQL数据库支持该设置，测试连接之后修改调用时必填。  **约束限制**：  - 测试连接之后修改调用时必填。  **取值范围**：  - true：迁移后，所有源数据库对象的Definer都会迁移至该用户下，其他用户需要授权后才具有数据库对象权限 - false：迁移后，将保持源数据库对象Definer定义不变，选择此选项，需要配合下一步用户权限迁移功能，将源数据库的用户全部迁移，这样才能保持源数据库的权限体系完全不变  **默认取值**：  不涉及
	ReplaceDefiner *bool `json:"replace_definer,omitempty"`

	// **参数解释**：  标签信息，用于给任务添加标识以便分类管理，详情请参见ResourceTag对象定义。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Tags *[]ResourceTag `json:"tags,omitempty"`

	// **参数解释**：  迁移任务的业务场景类型。  **约束限制**：  不涉及。  **取值范围**：  - migration：实时迁移 - sync：实时同步 - cloudDataGuard：实时灾备  **默认取值**：  不涉及
	DbUseType *ModifyJobReqDbUseType `json:"db_use_type,omitempty"`

	// **参数解释**：  产品ID，用于标识包周期任务的产品信息。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	ProductId *string `json:"product_id,omitempty"`
}

func (o ModifyJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyJobReq struct{}"
	}

	return strings.Join([]string{"ModifyJobReq", string(data)}, " ")
}

type ModifyJobReqTaskType struct {
	value string
}

type ModifyJobReqTaskTypeEnum struct {
	FULL_TRANS      ModifyJobReqTaskType
	INCR_TRANS      ModifyJobReqTaskType
	FULL_INCR_TRANS ModifyJobReqTaskType
}

func GetModifyJobReqTaskTypeEnum() ModifyJobReqTaskTypeEnum {
	return ModifyJobReqTaskTypeEnum{
		FULL_TRANS: ModifyJobReqTaskType{
			value: "FULL_TRANS",
		},
		INCR_TRANS: ModifyJobReqTaskType{
			value: "INCR_TRANS",
		},
		FULL_INCR_TRANS: ModifyJobReqTaskType{
			value: "FULL_INCR_TRANS",
		},
	}
}

func (c ModifyJobReqTaskType) Value() string {
	return c.value
}

func (c ModifyJobReqTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyJobReqTaskType) UnmarshalJSON(b []byte) error {
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

type ModifyJobReqNodeType struct {
	value string
}

type ModifyJobReqNodeTypeEnum struct {
	MICRO     ModifyJobReqNodeType
	SMALL     ModifyJobReqNodeType
	MEDIUM    ModifyJobReqNodeType
	HIGH      ModifyJobReqNodeType
	XLARGE    ModifyJobReqNodeType
	E_2XLARGE ModifyJobReqNodeType
}

func GetModifyJobReqNodeTypeEnum() ModifyJobReqNodeTypeEnum {
	return ModifyJobReqNodeTypeEnum{
		MICRO: ModifyJobReqNodeType{
			value: "micro",
		},
		SMALL: ModifyJobReqNodeType{
			value: "small",
		},
		MEDIUM: ModifyJobReqNodeType{
			value: "medium",
		},
		HIGH: ModifyJobReqNodeType{
			value: "high",
		},
		XLARGE: ModifyJobReqNodeType{
			value: "xlarge",
		},
		E_2XLARGE: ModifyJobReqNodeType{
			value: "2xlarge",
		},
	}
}

func (c ModifyJobReqNodeType) Value() string {
	return c.value
}

func (c ModifyJobReqNodeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyJobReqNodeType) UnmarshalJSON(b []byte) error {
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

type ModifyJobReqEngineType struct {
	value string
}

type ModifyJobReqEngineTypeEnum struct {
	MYSQL                  ModifyJobReqEngineType
	MONGODB                ModifyJobReqEngineType
	CLOUD_DATA_GUARD_MYSQL ModifyJobReqEngineType
	GAUSSDBV5              ModifyJobReqEngineType
	POSTGRESQL             ModifyJobReqEngineType
	MYSQL_TO_KAFKA         ModifyJobReqEngineType
	TAURUS_TO_KAFKA        ModifyJobReqEngineType
	GAUSSDBV5HA_TO_KAFKA   ModifyJobReqEngineType
}

func GetModifyJobReqEngineTypeEnum() ModifyJobReqEngineTypeEnum {
	return ModifyJobReqEngineTypeEnum{
		MYSQL: ModifyJobReqEngineType{
			value: "mysql",
		},
		MONGODB: ModifyJobReqEngineType{
			value: "mongodb",
		},
		CLOUD_DATA_GUARD_MYSQL: ModifyJobReqEngineType{
			value: "cloudDataGuard-mysql",
		},
		GAUSSDBV5: ModifyJobReqEngineType{
			value: "gaussdbv5",
		},
		POSTGRESQL: ModifyJobReqEngineType{
			value: "postgresql",
		},
		MYSQL_TO_KAFKA: ModifyJobReqEngineType{
			value: "mysql-to-kafka",
		},
		TAURUS_TO_KAFKA: ModifyJobReqEngineType{
			value: "taurus-to-kafka",
		},
		GAUSSDBV5HA_TO_KAFKA: ModifyJobReqEngineType{
			value: "gaussdbv5ha-to-kafka",
		},
	}
}

func (c ModifyJobReqEngineType) Value() string {
	return c.value
}

func (c ModifyJobReqEngineType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyJobReqEngineType) UnmarshalJSON(b []byte) error {
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

type ModifyJobReqNetType struct {
	value string
}

type ModifyJobReqNetTypeEnum struct {
	VPC ModifyJobReqNetType
	VPN ModifyJobReqNetType
	EIP ModifyJobReqNetType
}

func GetModifyJobReqNetTypeEnum() ModifyJobReqNetTypeEnum {
	return ModifyJobReqNetTypeEnum{
		VPC: ModifyJobReqNetType{
			value: "vpc",
		},
		VPN: ModifyJobReqNetType{
			value: "vpn",
		},
		EIP: ModifyJobReqNetType{
			value: "eip",
		},
	}
}

func (c ModifyJobReqNetType) Value() string {
	return c.value
}

func (c ModifyJobReqNetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyJobReqNetType) UnmarshalJSON(b []byte) error {
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

type ModifyJobReqJobDirection struct {
	value string
}

type ModifyJobReqJobDirectionEnum struct {
	UP      ModifyJobReqJobDirection
	DOWN    ModifyJobReqJobDirection
	NON_DBS ModifyJobReqJobDirection
}

func GetModifyJobReqJobDirectionEnum() ModifyJobReqJobDirectionEnum {
	return ModifyJobReqJobDirectionEnum{
		UP: ModifyJobReqJobDirection{
			value: "up",
		},
		DOWN: ModifyJobReqJobDirection{
			value: "down",
		},
		NON_DBS: ModifyJobReqJobDirection{
			value: "non-dbs",
		},
	}
}

func (c ModifyJobReqJobDirection) Value() string {
	return c.value
}

func (c ModifyJobReqJobDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyJobReqJobDirection) UnmarshalJSON(b []byte) error {
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

type ModifyJobReqDbUseType struct {
	value string
}

type ModifyJobReqDbUseTypeEnum struct {
	MIGRATION        ModifyJobReqDbUseType
	SYNC             ModifyJobReqDbUseType
	CLOUD_DATA_GUARD ModifyJobReqDbUseType
}

func GetModifyJobReqDbUseTypeEnum() ModifyJobReqDbUseTypeEnum {
	return ModifyJobReqDbUseTypeEnum{
		MIGRATION: ModifyJobReqDbUseType{
			value: "migration",
		},
		SYNC: ModifyJobReqDbUseType{
			value: "sync",
		},
		CLOUD_DATA_GUARD: ModifyJobReqDbUseType{
			value: "cloudDataGuard",
		},
	}
}

func (c ModifyJobReqDbUseType) Value() string {
	return c.value
}

func (c ModifyJobReqDbUseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyJobReqDbUseType) UnmarshalJSON(b []byte) error {
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
