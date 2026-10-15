package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateJobReq 创建任务请求体
type CreateJobReq struct {

	// **参数解释**：  是否绑定EIP（弹性公网IP）。  **约束限制**：  网络类型为EIP时必填且为true。  **取值范围**：  - true：绑定EIP - false：不绑定EIP  **默认取值**：  不涉及
	BindEip *bool `json:"bind_eip,omitempty"`

	// **参数解释**：  迁移任务的业务场景类型。  **约束限制**：  不涉及。  **取值范围**：  - migration：实时迁移 - sync：实时同步 - cloudDataGuard：实时灾备  **默认取值**：  不涉及
	DbUseType CreateJobReqDbUseType `json:"db_use_type"`

	// **参数解释**：  任务名称。  **约束限制**：  - 必须以字母开头，不区分大小写，可以包含字母、数字、中划线或下划线，不能包括其他特殊字符。  **取值范围**：  长度为4-50个字符。  **默认取值**：  不涉及
	Name string `json:"name"`

	// **参数解释**：  任务描述。  **约束限制**：  不能超过256位，且不能包含!<>&'\"特殊字符。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**：  DRS任务引擎类型。  **约束限制**：  不涉及。  **取值范围**：  - mysql：MySQL到MySQL迁移，MySQL到MySQL同步 - mongodb：MongoDB到DDS迁移与同步 - cloudDataGuard-mysql：MySQL到MySQL灾备 - gaussdbv5：GaussDB同步 - postgresql：PostgreSQL到PostgreSQL同步 - mysql-to-kafka：MySQL到Kafka同步 - taurus-to-kafka：TaurusDB到Kafka同步 - gaussdbv5ha-to-kafka：GaussDB集中式版到Kafka同步 - oracle-to-gaussdbv5：Oracle-GaussDB分布式版 - oracle-to-gaussdbv5ha：Oracle-GaussDB集中式版 - gaussdbv5-to-oracle：GaussDB分布式版-Oracle - gaussdbv5ha-to-oracle：GaussDB集中式版-Oracle - gaussdbv5-to-kafka：GaussDB分布式版-Kafka  详细取值可参考引擎类型说明。  **默认取值**：  不涉及
	EngineType CreateJobReqEngineType `json:"engine_type"`

	// **参数解释**：  指定目标实例是否限制为只读。  **约束限制**：  - MySQL迁移和灾备，且job_direction为up时设置有效。 - 灾备场景下，单主灾备且本云为备为必填且为true，不填默认设置为true。  **取值范围**：  - true：目标实例限制为只读 - false：目标实例不限制为只读  **默认取值**：  不涉及
	IsTargetReadonly *bool `json:"is_target_readonly,omitempty"`

	// **参数解释**：  任务的迁移方向。  **约束限制**：  不涉及。  **取值范围**：  - up：入云，灾备场景时对应本云为备 - down：出云，灾备场景时对应本云为主 - non-dbs：自建  **默认取值**：  不涉及
	JobDirection CreateJobReqJobDirection `json:"job_direction"`

	// **参数解释**：  是否为双主灾备模式。  **约束限制**：  - db_use_type为cloudDataGuard时必填。 - db_use_type为其他类型时为非必选参数。  **取值范围**：  - true：双主灾备 - false：非双主灾备  **默认取值**：  false
	MultiWrite *bool `json:"multi_write,omitempty"`

	// **参数解释**：  任务使用的网络类型。  **约束限制**：  灾备场景不支持选择VPC网络。  **取值范围**：  - vpn：虚拟专用网络 - vpc：虚拟私有云 - eip：弹性公网IP  **默认取值**：  不涉及
	NetType CreateJobReqNetType `json:"net_type"`

	// **参数解释**：  节点个数。MongoDB数据库时对应源端分片个数，源库为集群时必填。MySQL双主灾备时会默认设置为2。  **约束限制**：  不涉及。  **取值范围**：  1-32  **默认取值**：  不涉及
	NodeNum *int32 `json:"node_num,omitempty"`

	// **参数解释**：  DRS任务实例规格编码，决定迁移实例的计算资源大小。  **约束限制**：  不涉及。  **取值范围**：  - micro：极小规格 - small：小规格 - medium：中规格 - high：大规格 - xlarge：超大规格 - 2xlarge：极大规格  具体某种场景支持的取值可以通过查询可用的Node规格接口获取。  **默认取值**：  high
	NodeType CreateJobReqNodeType `json:"node_type"`

	SourceEndpoint *Endpoint `json:"source_endpoint"`

	TargetEndpoint *Endpoint `json:"target_endpoint"`

	// **参数解释**：  标签信息列表，详情请参见ResourceTag对象定义。  **约束限制**：  最多添加20个标签。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Tags *[]ResourceTag `json:"tags,omitempty"`

	// **参数解释**：  任务迁移模式，指定数据复制任务的执行方式。  **约束限制**：  单主灾备场景仅支持选择全量加增量（FULL_INCR_TRANS）。  **取值范围**：  - FULL_TRANS：全量迁移，仅迁移历史数据 - FULL_INCR_TRANS：全量+增量迁移，先迁移历史数据再持续同步增量 - INCR_TRANS：增量迁移，仅同步数据变更  **默认取值**：  FULL_INCR_TRANS
	TaskType CreateJobReqTaskType `json:"task_type"`

	// **参数解释**：  DRS实例所在子网ID，对应目标库相同VPC下已创建的子网（subnet）的网络ID。  **约束限制**：  UUID格式。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	CustomizeSutnetId string `json:"customize_sutnet_id"`

	// **参数解释**：  产品ID，用于标识包年包月任务对应的产品。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	ProductId *string `json:"product_id,omitempty"`

	// **参数解释**：  企业项目标签列表，详情请参见ResourceTag对象定义。  **约束限制**：  - 不填默认为default。 - key值必须为_sys_enterprise_project_id，value为企业项目ID。 - 只能填一个企业项目。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	SysTags *[]ResourceTag `json:"sys_tags,omitempty"`

	// **参数解释**：  任务处于异常状态一段时间后自动结束的天数。  **约束限制**：  不涉及。  **取值范围**：  14-100  **默认取值**：  14
	ExpiredDays *string `json:"expired_days,omitempty"`

	// **参数解释**：  主备任务主任务所在可用区编码，可以通过查询规格未售罄的可用区接口获取。  **约束限制**：  - master_az和slave_az同时填写时生效。 - 目前支持mysql，gaussdbv5ha-to-kafka。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	MasterAz *string `json:"master_az,omitempty"`

	// **参数解释**：  主备任务备任务所在可用区编码，可以通过查询规格未售罄的可用区接口获取。  **约束限制**：  - master_az和slave_az同时填写时生效。 - 目前支持mysql，gaussdbv5ha-to-kafka。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	SlaveAz *string `json:"slave_az,omitempty"`

	// **参数解释**：  任务的计费模式，指定数据复制任务的付费方式。  **约束限制**：  - db_use_type为migration时不支持包年包月。 - 支持的engine_type为：mysql，cloudDataGuard-mysql，postgresql，gaussdbv5ha-to-kafka。  **取值范围**：  - period：包年包月 - on_demand：按需计费  **默认取值**：  on_demand
	ChargingMode *CreateJobReqChargingMode `json:"charging_mode,omitempty"`

	PeriodOrder *PeriodOrderInfo `json:"period_order,omitempty"`

	// **参数解释**：  指定公网IP的信息列表，详情请参见PublicIpConfig对象定义。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	PublicIpList *[]PublicIpConfig `json:"public_ip_list,omitempty"`

	// **参数解释**：  是否开启云数据库RDS for MySQL或MariaDB的Binlog快速清理。  **约束限制**：  不涉及。  **取值范围**：  - true：开启Binlog快速清理 - false：不开启Binlog快速清理  **默认取值**：  false
	IsOpenFastClean *bool `json:"is_open_fast_clean,omitempty"`
}

func (o CreateJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateJobReq struct{}"
	}

	return strings.Join([]string{"CreateJobReq", string(data)}, " ")
}

type CreateJobReqDbUseType struct {
	value string
}

type CreateJobReqDbUseTypeEnum struct {
	MIGRATION        CreateJobReqDbUseType
	SYNC             CreateJobReqDbUseType
	CLOUD_DATA_GUARD CreateJobReqDbUseType
}

func GetCreateJobReqDbUseTypeEnum() CreateJobReqDbUseTypeEnum {
	return CreateJobReqDbUseTypeEnum{
		MIGRATION: CreateJobReqDbUseType{
			value: "migration",
		},
		SYNC: CreateJobReqDbUseType{
			value: "sync",
		},
		CLOUD_DATA_GUARD: CreateJobReqDbUseType{
			value: "cloudDataGuard",
		},
	}
}

func (c CreateJobReqDbUseType) Value() string {
	return c.value
}

func (c CreateJobReqDbUseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateJobReqDbUseType) UnmarshalJSON(b []byte) error {
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

type CreateJobReqEngineType struct {
	value string
}

type CreateJobReqEngineTypeEnum struct {
	MYSQL                  CreateJobReqEngineType
	MONGODB                CreateJobReqEngineType
	CLOUD_DATA_GUARD_MYSQL CreateJobReqEngineType
	GAUSSDBV5              CreateJobReqEngineType
	POSTGRESQL             CreateJobReqEngineType
	MYSQL_TO_KAFKA         CreateJobReqEngineType
	TAURUS_TO_KAFKA        CreateJobReqEngineType
	GAUSSDBV5HA_TO_KAFKA   CreateJobReqEngineType
}

func GetCreateJobReqEngineTypeEnum() CreateJobReqEngineTypeEnum {
	return CreateJobReqEngineTypeEnum{
		MYSQL: CreateJobReqEngineType{
			value: "mysql",
		},
		MONGODB: CreateJobReqEngineType{
			value: "mongodb",
		},
		CLOUD_DATA_GUARD_MYSQL: CreateJobReqEngineType{
			value: "cloudDataGuard-mysql",
		},
		GAUSSDBV5: CreateJobReqEngineType{
			value: "gaussdbv5",
		},
		POSTGRESQL: CreateJobReqEngineType{
			value: "postgresql",
		},
		MYSQL_TO_KAFKA: CreateJobReqEngineType{
			value: "mysql-to-kafka",
		},
		TAURUS_TO_KAFKA: CreateJobReqEngineType{
			value: "taurus-to-kafka",
		},
		GAUSSDBV5HA_TO_KAFKA: CreateJobReqEngineType{
			value: "gaussdbv5ha-to-kafka",
		},
	}
}

func (c CreateJobReqEngineType) Value() string {
	return c.value
}

func (c CreateJobReqEngineType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateJobReqEngineType) UnmarshalJSON(b []byte) error {
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

type CreateJobReqJobDirection struct {
	value string
}

type CreateJobReqJobDirectionEnum struct {
	UP      CreateJobReqJobDirection
	DOWN    CreateJobReqJobDirection
	NON_DBS CreateJobReqJobDirection
}

func GetCreateJobReqJobDirectionEnum() CreateJobReqJobDirectionEnum {
	return CreateJobReqJobDirectionEnum{
		UP: CreateJobReqJobDirection{
			value: "up",
		},
		DOWN: CreateJobReqJobDirection{
			value: "down",
		},
		NON_DBS: CreateJobReqJobDirection{
			value: "non-dbs",
		},
	}
}

func (c CreateJobReqJobDirection) Value() string {
	return c.value
}

func (c CreateJobReqJobDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateJobReqJobDirection) UnmarshalJSON(b []byte) error {
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

type CreateJobReqNetType struct {
	value string
}

type CreateJobReqNetTypeEnum struct {
	VPN CreateJobReqNetType
	VPC CreateJobReqNetType
	EIP CreateJobReqNetType
}

func GetCreateJobReqNetTypeEnum() CreateJobReqNetTypeEnum {
	return CreateJobReqNetTypeEnum{
		VPN: CreateJobReqNetType{
			value: "vpn",
		},
		VPC: CreateJobReqNetType{
			value: "vpc",
		},
		EIP: CreateJobReqNetType{
			value: "eip",
		},
	}
}

func (c CreateJobReqNetType) Value() string {
	return c.value
}

func (c CreateJobReqNetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateJobReqNetType) UnmarshalJSON(b []byte) error {
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

type CreateJobReqNodeType struct {
	value string
}

type CreateJobReqNodeTypeEnum struct {
	MICRO     CreateJobReqNodeType
	SMALL     CreateJobReqNodeType
	MEDIUM    CreateJobReqNodeType
	HIGH      CreateJobReqNodeType
	XLARGE    CreateJobReqNodeType
	E_2XLARGE CreateJobReqNodeType
}

func GetCreateJobReqNodeTypeEnum() CreateJobReqNodeTypeEnum {
	return CreateJobReqNodeTypeEnum{
		MICRO: CreateJobReqNodeType{
			value: "micro",
		},
		SMALL: CreateJobReqNodeType{
			value: "small",
		},
		MEDIUM: CreateJobReqNodeType{
			value: "medium",
		},
		HIGH: CreateJobReqNodeType{
			value: "high",
		},
		XLARGE: CreateJobReqNodeType{
			value: "xlarge",
		},
		E_2XLARGE: CreateJobReqNodeType{
			value: "2xlarge",
		},
	}
}

func (c CreateJobReqNodeType) Value() string {
	return c.value
}

func (c CreateJobReqNodeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateJobReqNodeType) UnmarshalJSON(b []byte) error {
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

type CreateJobReqTaskType struct {
	value string
}

type CreateJobReqTaskTypeEnum struct {
	FULL_TRANS      CreateJobReqTaskType
	FULL_INCR_TRANS CreateJobReqTaskType
	INCR_TRANS      CreateJobReqTaskType
}

func GetCreateJobReqTaskTypeEnum() CreateJobReqTaskTypeEnum {
	return CreateJobReqTaskTypeEnum{
		FULL_TRANS: CreateJobReqTaskType{
			value: "FULL_TRANS",
		},
		FULL_INCR_TRANS: CreateJobReqTaskType{
			value: "FULL_INCR_TRANS",
		},
		INCR_TRANS: CreateJobReqTaskType{
			value: "INCR_TRANS",
		},
	}
}

func (c CreateJobReqTaskType) Value() string {
	return c.value
}

func (c CreateJobReqTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateJobReqTaskType) UnmarshalJSON(b []byte) error {
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

type CreateJobReqChargingMode struct {
	value string
}

type CreateJobReqChargingModeEnum struct {
	PERIOD    CreateJobReqChargingMode
	ON_DEMAND CreateJobReqChargingMode
}

func GetCreateJobReqChargingModeEnum() CreateJobReqChargingModeEnum {
	return CreateJobReqChargingModeEnum{
		PERIOD: CreateJobReqChargingMode{
			value: "period",
		},
		ON_DEMAND: CreateJobReqChargingMode{
			value: "on_demand",
		},
	}
}

func (c CreateJobReqChargingMode) Value() string {
	return c.value
}

func (c CreateJobReqChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateJobReqChargingMode) UnmarshalJSON(b []byte) error {
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
