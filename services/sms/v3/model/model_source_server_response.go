package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SourceServerResponse 返回源端服务器信息
type SourceServerResponse struct {

	// 源端在SMS数据库中的ID
	Id *string `json:"id,omitempty"`

	// 源端服务器ip，注册源端时必选，更新非必选
	Ip string `json:"ip"`

	// 用来区分不同源端服务器的名称
	Name string `json:"name"`

	// 源端服务器的OS类型，分为Windows和Linux，注册必选，更新非必选
	OsType SourceServerResponseOsType `json:"os_type"`

	// 操作系统版本，注册必选，更新非必选
	OsVersion *string `json:"os_version,omitempty"`

	// 是否是OEM操作系统(Windows)
	OemSystem *bool `json:"oem_system,omitempty"`

	// 当前源端服务器状态 unavailable：环境校验不通过 waiting：等待 initialize：初始化 replicate：复制 syncing：持续同步 stopping：暂停中 stopped：已暂停 deleting：删除中 error：错误 cloning：等待克隆完成 testing：测试中 finished：启动目的端完成
	State *SourceServerResponseState `json:"state,omitempty"`

	// 迁移周期 cutovering:启动目的端中 cutovered:启动目的端完成 checking:检查中 setting:设置中 replicating:复制中 syncing:同步中
	MigrationCycle *SourceServerResponseMigrationCycle `json:"migration_cycle,omitempty"`
}

func (o SourceServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourceServerResponse struct{}"
	}

	return strings.Join([]string{"SourceServerResponse", string(data)}, " ")
}

type SourceServerResponseOsType struct {
	value string
}

type SourceServerResponseOsTypeEnum struct {
	WINDOWS SourceServerResponseOsType
	LINUX   SourceServerResponseOsType
}

func GetSourceServerResponseOsTypeEnum() SourceServerResponseOsTypeEnum {
	return SourceServerResponseOsTypeEnum{
		WINDOWS: SourceServerResponseOsType{
			value: "WINDOWS",
		},
		LINUX: SourceServerResponseOsType{
			value: "LINUX",
		},
	}
}

func (c SourceServerResponseOsType) Value() string {
	return c.value
}

func (c SourceServerResponseOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceServerResponseOsType) UnmarshalJSON(b []byte) error {
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

type SourceServerResponseState struct {
	value string
}

type SourceServerResponseStateEnum struct {
	UNAVAILABLE SourceServerResponseState
	WAITING     SourceServerResponseState
	INITIALIZE  SourceServerResponseState
	REPLICATE   SourceServerResponseState
	SYNCING     SourceServerResponseState
	STOPPING    SourceServerResponseState
	STOPPED     SourceServerResponseState
	DELETING    SourceServerResponseState
	ERROR       SourceServerResponseState
	CLONING     SourceServerResponseState
	TESTING     SourceServerResponseState
	FINISHED    SourceServerResponseState
}

func GetSourceServerResponseStateEnum() SourceServerResponseStateEnum {
	return SourceServerResponseStateEnum{
		UNAVAILABLE: SourceServerResponseState{
			value: "unavailable",
		},
		WAITING: SourceServerResponseState{
			value: "waiting",
		},
		INITIALIZE: SourceServerResponseState{
			value: "initialize",
		},
		REPLICATE: SourceServerResponseState{
			value: "replicate",
		},
		SYNCING: SourceServerResponseState{
			value: "syncing",
		},
		STOPPING: SourceServerResponseState{
			value: "stopping",
		},
		STOPPED: SourceServerResponseState{
			value: "stopped",
		},
		DELETING: SourceServerResponseState{
			value: "deleting",
		},
		ERROR: SourceServerResponseState{
			value: "error",
		},
		CLONING: SourceServerResponseState{
			value: "cloning",
		},
		TESTING: SourceServerResponseState{
			value: "testing",
		},
		FINISHED: SourceServerResponseState{
			value: "finished",
		},
	}
}

func (c SourceServerResponseState) Value() string {
	return c.value
}

func (c SourceServerResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceServerResponseState) UnmarshalJSON(b []byte) error {
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

type SourceServerResponseMigrationCycle struct {
	value string
}

type SourceServerResponseMigrationCycleEnum struct {
	CUTOVERING  SourceServerResponseMigrationCycle
	CUTOVERED   SourceServerResponseMigrationCycle
	CHECKING    SourceServerResponseMigrationCycle
	SETTING     SourceServerResponseMigrationCycle
	REPLICATING SourceServerResponseMigrationCycle
	SYNCING     SourceServerResponseMigrationCycle
}

func GetSourceServerResponseMigrationCycleEnum() SourceServerResponseMigrationCycleEnum {
	return SourceServerResponseMigrationCycleEnum{
		CUTOVERING: SourceServerResponseMigrationCycle{
			value: "cutovering",
		},
		CUTOVERED: SourceServerResponseMigrationCycle{
			value: "cutovered",
		},
		CHECKING: SourceServerResponseMigrationCycle{
			value: "checking",
		},
		SETTING: SourceServerResponseMigrationCycle{
			value: "setting",
		},
		REPLICATING: SourceServerResponseMigrationCycle{
			value: "replicating",
		},
		SYNCING: SourceServerResponseMigrationCycle{
			value: "syncing",
		},
	}
}

func (c SourceServerResponseMigrationCycle) Value() string {
	return c.value
}

func (c SourceServerResponseMigrationCycle) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceServerResponseMigrationCycle) UnmarshalJSON(b []byte) error {
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
