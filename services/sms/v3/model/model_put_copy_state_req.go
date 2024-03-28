package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PutCopyStateReq 源端复制状态
type PutCopyStateReq struct {

	// 源端服务器状 UNAVAILABLE：环境校验不通过 WAITING：等待 INIT：初始化 REPLICATE：复制 SYNCING：持续同步 STOPPING：暂停中 STOPPED：已暂停 DELETING：删除中 ERROR：错误 CLONING：等待克隆完成 CUTOVERING：启动目的端中 FINISHED：启动目的端完成
	Copystate *PutCopyStateReqCopystate `json:"copystate,omitempty"`

	// 迁移周期 cutovering:启动目的端中 cutovered:启动目的端完成 checking:检查中 setting:设置中 replicating:复制中 syncing:同步中
	Migrationcycle *PutCopyStateReqMigrationcycle `json:"migrationcycle,omitempty"`
}

func (o PutCopyStateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutCopyStateReq struct{}"
	}

	return strings.Join([]string{"PutCopyStateReq", string(data)}, " ")
}

type PutCopyStateReqCopystate struct {
	value string
}

type PutCopyStateReqCopystateEnum struct {
	UNAVAILABLE PutCopyStateReqCopystate
	WAITING     PutCopyStateReqCopystate
	INIT        PutCopyStateReqCopystate
	REPLICATE   PutCopyStateReqCopystate
	SYNCING     PutCopyStateReqCopystate
	STOPPING    PutCopyStateReqCopystate
	STOPPED     PutCopyStateReqCopystate
	DELETING    PutCopyStateReqCopystate
	ERROR       PutCopyStateReqCopystate
	CLONING     PutCopyStateReqCopystate
	CUTOVERING  PutCopyStateReqCopystate
	FINISHED    PutCopyStateReqCopystate
}

func GetPutCopyStateReqCopystateEnum() PutCopyStateReqCopystateEnum {
	return PutCopyStateReqCopystateEnum{
		UNAVAILABLE: PutCopyStateReqCopystate{
			value: "UNAVAILABLE",
		},
		WAITING: PutCopyStateReqCopystate{
			value: "WAITING",
		},
		INIT: PutCopyStateReqCopystate{
			value: "INIT",
		},
		REPLICATE: PutCopyStateReqCopystate{
			value: "REPLICATE",
		},
		SYNCING: PutCopyStateReqCopystate{
			value: "SYNCING",
		},
		STOPPING: PutCopyStateReqCopystate{
			value: "STOPPING",
		},
		STOPPED: PutCopyStateReqCopystate{
			value: "STOPPED",
		},
		DELETING: PutCopyStateReqCopystate{
			value: "DELETING",
		},
		ERROR: PutCopyStateReqCopystate{
			value: "ERROR",
		},
		CLONING: PutCopyStateReqCopystate{
			value: "CLONING",
		},
		CUTOVERING: PutCopyStateReqCopystate{
			value: "CUTOVERING",
		},
		FINISHED: PutCopyStateReqCopystate{
			value: "FINISHED",
		},
	}
}

func (c PutCopyStateReqCopystate) Value() string {
	return c.value
}

func (c PutCopyStateReqCopystate) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PutCopyStateReqCopystate) UnmarshalJSON(b []byte) error {
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

type PutCopyStateReqMigrationcycle struct {
	value string
}

type PutCopyStateReqMigrationcycleEnum struct {
	CUTOVERING  PutCopyStateReqMigrationcycle
	CUTOVERED   PutCopyStateReqMigrationcycle
	CHECKING    PutCopyStateReqMigrationcycle
	SETTING     PutCopyStateReqMigrationcycle
	REPLICATING PutCopyStateReqMigrationcycle
	SYNCING     PutCopyStateReqMigrationcycle
}

func GetPutCopyStateReqMigrationcycleEnum() PutCopyStateReqMigrationcycleEnum {
	return PutCopyStateReqMigrationcycleEnum{
		CUTOVERING: PutCopyStateReqMigrationcycle{
			value: "cutovering",
		},
		CUTOVERED: PutCopyStateReqMigrationcycle{
			value: "cutovered",
		},
		CHECKING: PutCopyStateReqMigrationcycle{
			value: "checking",
		},
		SETTING: PutCopyStateReqMigrationcycle{
			value: "setting",
		},
		REPLICATING: PutCopyStateReqMigrationcycle{
			value: "replicating",
		},
		SYNCING: PutCopyStateReqMigrationcycle{
			value: "syncing",
		},
	}
}

func (c PutCopyStateReqMigrationcycle) Value() string {
	return c.value
}

func (c PutCopyStateReqMigrationcycle) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PutCopyStateReqMigrationcycle) UnmarshalJSON(b []byte) error {
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
