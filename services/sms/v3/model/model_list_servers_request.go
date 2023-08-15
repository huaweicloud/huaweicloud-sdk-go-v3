package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListServersRequest Request Object
type ListServersRequest struct {

	// 源端服务器状态 unavailable：环境校验不通过 waiting：等待 initialize：初始化 replicate：复制 syncing：持续同步 stopping：暂停中 stopped：已暂停 deleting：删除中 error：错误 cloning：等待克隆完成 cutovering：启动目的端中 finished：启动目的端完成
	State *ListServersRequestState `json:"state,omitempty"`

	// 源端服务器名称
	Name *string `json:"name,omitempty"`

	// 源端服务器ID
	Id *string `json:"id,omitempty"`

	// 源端服务器IP地址
	Ip *string `json:"ip,omitempty"`

	// 迁移项目ID，填写该参数将查询迁移项目下的所有虚拟机
	Migproject *string `json:"migproject,omitempty"`

	// 每一页记录的源端服务器数量，0表示用默认值 200
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，默认值0
	Offset *int32 `json:"offset,omitempty"`

	// checking:检查中 setting:设置中 replicating:复制中 syncing:同步中 cutovering:启动目的端中 cutovered:启动目的端完成
	MigrationCycle *ListServersRequestMigrationCycle `json:"migration_cycle,omitempty"`

	// 查询失去连接的源端
	Connected *bool `json:"connected,omitempty"`

	// 需要查询的企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServersRequest struct{}"
	}

	return strings.Join([]string{"ListServersRequest", string(data)}, " ")
}

type ListServersRequestState struct {
	value string
}

type ListServersRequestStateEnum struct {
	UNAVAILABLE ListServersRequestState
	WAITING     ListServersRequestState
	INITIALIZE  ListServersRequestState
	REPLICATE   ListServersRequestState
	SYNCING     ListServersRequestState
	STOPPING    ListServersRequestState
	STOPPED     ListServersRequestState
	DELETING    ListServersRequestState
	ERROR       ListServersRequestState
	CLONING     ListServersRequestState
	CUTOVERING  ListServersRequestState
	FINISHED    ListServersRequestState
}

func GetListServersRequestStateEnum() ListServersRequestStateEnum {
	return ListServersRequestStateEnum{
		UNAVAILABLE: ListServersRequestState{
			value: "unavailable",
		},
		WAITING: ListServersRequestState{
			value: "waiting",
		},
		INITIALIZE: ListServersRequestState{
			value: "initialize",
		},
		REPLICATE: ListServersRequestState{
			value: "replicate",
		},
		SYNCING: ListServersRequestState{
			value: "syncing",
		},
		STOPPING: ListServersRequestState{
			value: "stopping",
		},
		STOPPED: ListServersRequestState{
			value: "stopped",
		},
		DELETING: ListServersRequestState{
			value: "deleting",
		},
		ERROR: ListServersRequestState{
			value: "error",
		},
		CLONING: ListServersRequestState{
			value: "cloning",
		},
		CUTOVERING: ListServersRequestState{
			value: "cutovering",
		},
		FINISHED: ListServersRequestState{
			value: "finished",
		},
	}
}

func (c ListServersRequestState) Value() string {
	return c.value
}

func (c ListServersRequestState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListServersRequestState) UnmarshalJSON(b []byte) error {
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

type ListServersRequestMigrationCycle struct {
	value string
}

type ListServersRequestMigrationCycleEnum struct {
	CHECKING    ListServersRequestMigrationCycle
	SETTING     ListServersRequestMigrationCycle
	REPLICATING ListServersRequestMigrationCycle
	SYNCING     ListServersRequestMigrationCycle
	CUTOVERING  ListServersRequestMigrationCycle
	CUTOVERED   ListServersRequestMigrationCycle
}

func GetListServersRequestMigrationCycleEnum() ListServersRequestMigrationCycleEnum {
	return ListServersRequestMigrationCycleEnum{
		CHECKING: ListServersRequestMigrationCycle{
			value: "checking",
		},
		SETTING: ListServersRequestMigrationCycle{
			value: "setting",
		},
		REPLICATING: ListServersRequestMigrationCycle{
			value: "replicating",
		},
		SYNCING: ListServersRequestMigrationCycle{
			value: "syncing",
		},
		CUTOVERING: ListServersRequestMigrationCycle{
			value: "cutovering",
		},
		CUTOVERED: ListServersRequestMigrationCycle{
			value: "cutovered",
		},
	}
}

func (c ListServersRequestMigrationCycle) Value() string {
	return c.value
}

func (c ListServersRequestMigrationCycle) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListServersRequestMigrationCycle) UnmarshalJSON(b []byte) error {
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
