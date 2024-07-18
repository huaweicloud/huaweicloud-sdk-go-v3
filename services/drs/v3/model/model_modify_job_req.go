package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ModifyJobReq 修改任务描述信息、名称，设置异常通知信息、限速等。
type ModifyJobReq struct {

	// 任务id
	JobId string `json:"job_id"`

	// 任务描述，修改任务描述时必填。
	Description *string `json:"description,omitempty"`

	// 任务名称，修改任务名称时必填
	Name *string `json:"name,omitempty"`

	AlarmNotify *AlarmNotifyInfo `json:"alarm_notify,omitempty"`

	// 任务模式，FULL_TRANS：全量；FULL_INCR_TRANS：全量+增量；INCR_TRANS：增量。
	TaskType *ModifyJobReqTaskType `json:"task_type,omitempty"`

	SourceEndpoint *Endpoint `json:"source_endpoint,omitempty"`

	TargetEndpoint *Endpoint `json:"target_endpoint,omitempty"`

	// node规格类型，测试连接之后修改调用时必填。取值： - micro：极小规格。 - small：小规格。 - medium：中规格。 - high：大规格。 - xlarge：超大规格。 - 2xlarge：极大规格。
	NodeType *ModifyJobReqNodeType `json:"node_type,omitempty"`

	// 引擎类型，测试连接之后修改调用时必填。mysql：迁移，同步使用。mongodb：迁移使用。cloudDataGuard-mysql：灾备使用
	EngineType *ModifyJobReqEngineType `json:"engine_type,omitempty"`

	// 网络类型，测试连接之后修改调用时必填。
	NetType *ModifyJobReqNetType `json:"net_type,omitempty"`

	// 保存数据库信息，测试连接之后修改调用时必填为true。
	StoreDbInfo *bool `json:"store_db_info,omitempty"`

	// 是否为重建任务。
	IsRecreate *bool `json:"is_recreate,omitempty"`

	// 迁移方向,up 入云 灾备场景时对应本云为备,down 出云 灾备场景时对应本云为主,non-dbs 自建
	JobDirection *ModifyJobReqJobDirection `json:"job_direction,omitempty"`

	// 目标实例是否限制为只读。
	IsTargetReadonly *bool `json:"is_target_readonly,omitempty"`

	// 所有Definer是否迁移到该用户下，MySQL数据库支持该设置，测试连接之后修改调用时必填。 - true：迁移后，所有源数据库对象的Definer都会迁移至该用户下，其他用户需要授权后才具有数据库对象权限 - false：迁移后，将保持源数据库对象Definer定义不变，选择此选项，需要配合下一步用户权限迁移功能，将源数据库的用户全部迁移，这样才能保持源数据库的权限体系完全不变。
	ReplaceDefiner *bool `json:"replace_definer,omitempty"`

	// 标签信息
	Tags *[]ResourceTag `json:"tags,omitempty"`

	// 迁移类型，migration-实时迁移,sync-实时同步,cloudDataGuard-实时灾备
	DbUseType *ModifyJobReqDbUseType `json:"db_use_type,omitempty"`

	// 产品ID。
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
