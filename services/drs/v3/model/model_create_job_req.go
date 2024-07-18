package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateJobReq 创建任务请求体
type CreateJobReq struct {

	// 是否绑定eip，网络类型为eip时必填且为true
	BindEip *bool `json:"bind_eip,omitempty"`

	// 迁移场景，migration-实时迁移,sync-实时同步,cloudDataGuard-实时灾备
	DbUseType CreateJobReqDbUseType `json:"db_use_type"`

	// 任务名称，约束：任务名称在4位到50位之间，不区分大小写，可以包含字母、数字、中划线或下划线，不能包括其他特殊字符。
	Name string `json:"name"`

	// 任务描述。  **约束**：任务描述不能超过256位，且不能包含!<>&'\"\\特殊字符。
	Description *string `json:"description,omitempty"`

	// 引擎类型 - mysql：MySQL到MySQL迁移，MySQL到MySQL同步 - mongodb：MongoDB到DDS迁移 - cloudDataGuard-mysql：MySQL到MySQL灾备 - gaussdbv5：GaussDB同步 - mysql-to-kafka：MySQL到Kafka同步 - taurus-to-kafka：GaussDB(for MySQL)到Kafka同步 - gaussdbv5ha-to-kafka：GaussDB主备版到Kafka同步 - postgresql：PostgreSQL到PostgreSQL同步
	EngineType CreateJobReqEngineType `json:"engine_type"`

	// 指定目标实例是否限制为只读，MySQL迁移和灾备，且job_direction为up时设置有效。（灾备场景下，单主灾备且本云为备为必填且为true，不填默认设置为true）。
	IsTargetReadonly *bool `json:"is_target_readonly,omitempty"`

	// 迁移方向，up ：入云 ，灾备场景时对应本云为备，down：出云，灾备场景时对应本云为主，non-dbs：自建
	JobDirection CreateJobReqJobDirection `json:"job_direction"`

	// - db_use_type 是cloudDataGuard时，必填，灾备类型是双主灾备时 muti_write取值true, 否则为false。 - db_use_type 是其他类型时，muti_write是非必选参数
	MultiWrite *bool `json:"multi_write,omitempty"`

	// 网络类型
	NetType CreateJobReqNetType `json:"net_type"`

	// 节点个数。MongoDB数据库时对应源端分片个数，源库为集群时必填，[1-32]，MySQL双主灾备时会默认设置为2。
	NodeNum *int32 `json:"node_num,omitempty"`

	// 规格类型。取值： - micro：极小规格。 - small：小规格。 - medium：中规格。 - high：大规格。 - xlarge：超大规格。 - 2xlarge：极大规格。 具体某种场景支持的取值可以通过[查询可用的Node规格接口](https://support.huaweicloud.com/api-drs/drs_03_0239.html)获取。
	NodeType CreateJobReqNodeType `json:"node_type"`

	SourceEndpoint *Endpoint `json:"source_endpoint"`

	TargetEndpoint *Endpoint `json:"target_endpoint"`

	// 标签信息。
	Tags *[]ResourceTag `json:"tags,omitempty"`

	// 迁移模式，FULL_TRANS 全量,FULL_INCR_TRANS 全量+增量,INCR_TRANS 增量，灾备场景单主灾备仅支持全量加增量（FULL_INCR_TRANS）
	TaskType CreateJobReqTaskType `json:"task_type"`

	// DRS实例所在子网ID，对应目标库相同VPC下已创建的子网（subnet）的网络ID，UUID格式。
	CustomizeSutnetId string `json:"customize_sutnet_id"`

	// 产品id。
	ProductId *string `json:"product_id,omitempty"`

	// 企业项目，不填默认为default，key值必须为_sys_enterprise_project_id，value为企业项目ID，只能填一个企业项目。
	SysTags *[]ResourceTag `json:"sys_tags,omitempty"`

	// 任务处于异常状态一段时间后，将会自动结束，单位为天。(范围14-100)，不传默认为14天。
	ExpiredDays *string `json:"expired_days,omitempty"`

	// 主备任务主任务所在可用区code，可以通过查询规格未售罄的可用区接口获取 - master_az和slave_az同时填写时生效 - 目前支持mysql，gaussdbv5ha-to-kafka
	MasterAz *string `json:"master_az,omitempty"`

	// 主备任务备任务所在可用区code，可以通过查询规格未售罄的可用区接口获取 - master_az和slave_az同时填写时生效 - 目前支持mysql，gaussdbv5ha-to-kafka
	SlaveAz *string `json:"slave_az,omitempty"`

	// 计费模式，不填默认为按需计费。 取值范围： - period：包年包月。 - on_demand：按需计费。
	ChargingMode *CreateJobReqChargingMode `json:"charging_mode,omitempty"`

	PeriodOrder *PeriodOrderInfo `json:"period_order,omitempty"`

	// 指定公网IP的信息。
	PublicIpList *[]PublicIpConfig `json:"public_ip_list,omitempty"`

	// 是否开启云数据库RDS for MySQL/MariaDB的binlog快速清理。不传默认为false，不开启快速清理。
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
