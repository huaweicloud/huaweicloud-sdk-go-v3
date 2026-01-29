package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PutCopyStateReq 源端复制状态 copystate与migrationcycle忽略大小写，但必须满足一些搭配，否则更新失败，接口报错： migrationcycle状态为 checking时，copystate只能为unavailable/waiting setting时， waiting/premigready/premiging/premiged/premigfailed replicating时，init/replicate/deleting/stopping/error/stopped syncing时，syncing/cloning/deleting/stopping/error/stopped cutovering时，cutovering/deleting/stopping/error/stopped cutovered时，finished/deleting/stopping/error/stopped
type PutCopyStateReq struct {

	// 源端服务器状态 unavailable: 环境校验不通过 waiting: 等待 init: 初始化 replicate: 复制 syncing: 持续同步 stopping: 暂停中 stopped: 已暂停 skipping: 跳过中 deleting: 删除中 clearing: 清理快照资源中 cleared: 清理快照资源完成 clearfailed: 清理快照资源失败 premigready: 迁移演练就绪 premiging: 迁移演练中 premiged: 迁移演练完成 premigfailed: 迁移演练失败 cloning: 等待克隆完成 cutovering: 启动目的端中 finished: 启动目的端完成 error: 错误
	Copystate PutCopyStateReqCopystate `json:"copystate"`

	// 迁移周期 no_ready:未就绪 ready_for_test:已就绪 testing:测试中 tested:测试完成 cutovering:启动目的端中 cutovered:启动目的端完成 checking:检查中 setting:设置中 replicating:复制中 syncing:同步中
	Migrationcycle PutCopyStateReqMigrationcycle `json:"migrationcycle"`
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
	INIT         PutCopyStateReqCopystate
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
	PREMIGING    PutCopyStateReqCopystate
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
		INIT: PutCopyStateReqCopystate{
			value: "init",
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
		PREMIGING: PutCopyStateReqCopystate{
			value: "premiging",
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
	NO_READY       PutCopyStateReqMigrationcycle
	READY_FOR_TEST PutCopyStateReqMigrationcycle
	TESTING        PutCopyStateReqMigrationcycle
	TESTED         PutCopyStateReqMigrationcycle
	CUTOVERING     PutCopyStateReqMigrationcycle
	CUTOVERED      PutCopyStateReqMigrationcycle
	CHECKING       PutCopyStateReqMigrationcycle
	SETTING        PutCopyStateReqMigrationcycle
	REPLICATING    PutCopyStateReqMigrationcycle
	SYNCING        PutCopyStateReqMigrationcycle
}

func GetPutCopyStateReqMigrationcycleEnum() PutCopyStateReqMigrationcycleEnum {
	return PutCopyStateReqMigrationcycleEnum{
		NO_READY: PutCopyStateReqMigrationcycle{
			value: "no_ready",
		},
		READY_FOR_TEST: PutCopyStateReqMigrationcycle{
			value: "ready_for_test",
		},
		TESTING: PutCopyStateReqMigrationcycle{
			value: "testing",
		},
		TESTED: PutCopyStateReqMigrationcycle{
			value: "tested",
		},
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
