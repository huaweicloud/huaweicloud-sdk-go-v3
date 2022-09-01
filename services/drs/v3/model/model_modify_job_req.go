package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 修改任务描述信息、名称，设置异常通知信息、限速等。
type ModifyJobReq struct {

	// 任务id
	JobId string `json:"job_id" xml:"job_id"`

	// 任务描述，修改任务描述时必填。
	Description *string `json:"description,omitempty" xml:"description"`

	// 任务名称，修改任务名称时必填
	Name *string `json:"name,omitempty" xml:"name"`

	AlarmNotify *AlarmNotifyInfo `json:"alarm_notify,omitempty" xml:"alarm_notify"`

	// 任务模式，FULL_TRANS：全量；FULL_INCR_TRANS：全量+增量；INCR_TRANS：增量。
	TaskType *ModifyJobReqTaskType `json:"task_type,omitempty" xml:"task_type"`

	SourceEndpoint *Endpoint `json:"source_endpoint,omitempty" xml:"source_endpoint"`

	TargetEndpoint *Endpoint `json:"target_endpoint,omitempty" xml:"target_endpoint"`

	// node规格类型，测试连接之后修改调用时必填。
	NodeType *ModifyJobReqNodeType `json:"node_type,omitempty" xml:"node_type"`

	// 引擎类型，测试连接之后修改调用时必填。mysql：迁移，同步使用。mongodb：迁移使用。cloudDataGuard-mysql：灾备使用
	EngineType *ModifyJobReqEngineType `json:"engine_type,omitempty" xml:"engine_type"`

	// 网络类型，测试连接之后修改调用时必填。
	NetType *ModifyJobReqNetType `json:"net_type,omitempty" xml:"net_type"`

	// 保存数据库信息，测试连接之后修改调用时必填为true。
	StoreDbInfo *bool `json:"store_db_info,omitempty" xml:"store_db_info"`

	// 是否为重建任务。
	IsRecreate *bool `json:"is_recreate,omitempty" xml:"is_recreate"`

	// 迁移方向,up 入云 灾备场景时对应本云为备,down 出云 灾备场景时对应本云为主,non-dbs 自建
	JobDirection *ModifyJobReqJobDirection `json:"job_direction,omitempty" xml:"job_direction"`

	// 目标实例是否限制为只读。
	IsTargetReadonly *bool `json:"is_target_readonly,omitempty" xml:"is_target_readonly"`

	// 所有Definer是否迁移到该用户下，MySQL数据库支持该设置，测试连接之后修改调用时必填。 - true：迁移后，所有源数据库对象的Definer都会迁移至该用户下，其他用户需要授权后才具有数据库对象权限 - false：迁移后，将保持源数据库对象Definer定义不变，选择此选项，需要配合下一步用户权限迁移功能，将源数据库的用户全部迁移，这样才能保持源数据库的权限体系完全不变。
	ReplaceDefiner *bool `json:"replace_definer,omitempty" xml:"replace_definer"`

	// 标签信息
	Tags *[]ResourceTag `json:"tags,omitempty" xml:"tags"`

	// 迁移类型，migration-实时迁移,sync-实时同步,cloudDataGuard-实时灾备
	DbUseType *ModifyJobReqDbUseType `json:"db_use_type,omitempty" xml:"db_use_type"`

	// 产品ID。
	ProductId *string `json:"product_id,omitempty" xml:"product_id"`
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

type ModifyJobReqNodeType struct {
	value string
}

type ModifyJobReqNodeTypeEnum struct {
	HIGH ModifyJobReqNodeType
}

func GetModifyJobReqNodeTypeEnum() ModifyJobReqNodeTypeEnum {
	return ModifyJobReqNodeTypeEnum{
		HIGH: ModifyJobReqNodeType{
			value: "high",
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

type ModifyJobReqEngineType struct {
	value string
}

type ModifyJobReqEngineTypeEnum struct {
	MYSQL                  ModifyJobReqEngineType
	MONGODB                ModifyJobReqEngineType
	CLOUD_DATA_GUARD_MYSQL ModifyJobReqEngineType
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
