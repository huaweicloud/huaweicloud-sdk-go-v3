package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PutCopyStateReq 源端复制状态
type PutCopyStateReq struct {

	// 源端服务器状态 unavailable：环境校验不通过 waiting：等待 initialize：初始化 replicate：复制 syncing：持续同步 stopping：暂停中 stopped：已暂停 skipping：跳过中 deleting：删除中 clearing: 清理快照资源中 cleared：清理快照资源完成 clearfailed：清理快照资源失败 premigready：迁移演练就绪 premiged：迁移演练完成 premigfailed：迁移演练失败 cloning：等待克隆完成 cutovering：启动目的端中 finished：启动目的端完成 error：错误
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
	UNAVAILABLE  PutCopyStateReqCopystate
	WAITING      PutCopyStateReqCopystate
	INITIALIZE   PutCopyStateReqCopystate
	REPLICATE    PutCopyStateReqCopystate
	SYNCING      PutCopyStateReqCopystate
	STOPPING     PutCopyStateReqCopystate
	STOPPED      PutCopyStateReqCopystate
	SKIPPING     PutCopyStateReqCopystate
	DELETING     PutCopyStateReqCopystate
	CLEARING     PutCopyStateReqCopystate
	CLEARED      PutCopyStateReqCopystate
	CLEARFAILED  PutCopyStateReqCopystate
	PREMIGREADY  PutCopyStateReqCopystate
	PREMIGED     PutCopyStateReqCopystate
	PREMIGFAILED PutCopyStateReqCopystate
	CLONING      PutCopyStateReqCopystate
	CUTOVERING   PutCopyStateReqCopystate
	FINISHED     PutCopyStateReqCopystate
	ERROR        PutCopyStateReqCopystate
}

func GetPutCopyStateReqCopystateEnum() PutCopyStateReqCopystateEnum {
	return PutCopyStateReqCopystateEnum{
		UNAVAILABLE: PutCopyStateReqCopystate{
			value: "unavailable",
		},
		WAITING: PutCopyStateReqCopystate{
			value: "waiting",
		},
		INITIALIZE: PutCopyStateReqCopystate{
			value: "initialize",
		},
		REPLICATE: PutCopyStateReqCopystate{
			value: "replicate",
		},
		SYNCING: PutCopyStateReqCopystate{
			value: "syncing",
		},
		STOPPING: PutCopyStateReqCopystate{
			value: "stopping",
		},
		STOPPED: PutCopyStateReqCopystate{
			value: "stopped",
		},
		SKIPPING: PutCopyStateReqCopystate{
			value: "skipping",
		},
		DELETING: PutCopyStateReqCopystate{
			value: "deleting",
		},
		CLEARING: PutCopyStateReqCopystate{
			value: "clearing",
		},
		CLEARED: PutCopyStateReqCopystate{
			value: "cleared",
		},
		CLEARFAILED: PutCopyStateReqCopystate{
			value: "clearfailed",
		},
		PREMIGREADY: PutCopyStateReqCopystate{
			value: "premigready",
		},
		PREMIGED: PutCopyStateReqCopystate{
			value: "premiged",
		},
		PREMIGFAILED: PutCopyStateReqCopystate{
			value: "premigfailed",
		},
		CLONING: PutCopyStateReqCopystate{
			value: "cloning",
		},
		CUTOVERING: PutCopyStateReqCopystate{
			value: "cutovering",
		},
		FINISHED: PutCopyStateReqCopystate{
			value: "finished",
		},
		ERROR: PutCopyStateReqCopystate{
			value: "error",
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
