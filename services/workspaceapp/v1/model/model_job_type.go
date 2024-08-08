package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobType job类型： * `CREATE_SERVER` - 创建服务器 * `DELETE_SERVER` - 删除服务器 * `UPDATE_FREEZE_STATUS` - 修改服务器冻结状态 * `CREATE_SERVER_IMAGE` - 构建镜像 * `REINSTALL_OS` - 重装操作系统 * `CHANGE_SERVER_IMAGE` - 更换镜像 * `REJOIN_DOMAIN` - 重新加域 * `MIGRATE_SERVER` - 迁移服务器 * `UPGRADE_ACCESS_AGENT` - hda升级 * `UPDATE_SERVER_TSVI` - 更新虚拟会话IP * `SCHEDULED_TASK` - 定时任务job * `COLLECT_HDA_LOG` - 收集hda日志 * `COLLECT_APS_LOG` - 收集aps日志 * `CREATE_SERVER_SNAPSHOT` - 创建服务器快照 * `DELETE_SERVER_SNAPSHOT` - 删除服务器快照 * `RESTORE_SERVER_SNAPSHOT` - 恢复服务器快照
type JobType struct {
	value string
}

type JobTypeEnum struct {
	CREATE_SERVER           JobType
	DELETE_SERVER           JobType
	UPDATE_FREEZE_STATUS    JobType
	CREATE_SERVER_IMAGE     JobType
	REINSTALL_OS            JobType
	CHANGE_SERVER_IMAGE     JobType
	REJOIN_DOMAIN           JobType
	MIGRATE_SERVER          JobType
	UPGRADE_ACCESS_AGENT    JobType
	UPDATE_SERVER_TSVI      JobType
	SCHEDULED_TASK          JobType
	COLLECT_HDA_LOG         JobType
	COLLECT_APS_LOG         JobType
	CREATE_SERVER_SNAPSHOT  JobType
	DELETE_SERVER_SNAPSHOT  JobType
	RESTORE_SERVER_SNAPSHOT JobType
}

func GetJobTypeEnum() JobTypeEnum {
	return JobTypeEnum{
		CREATE_SERVER: JobType{
			value: "CREATE_SERVER",
		},
		DELETE_SERVER: JobType{
			value: "DELETE_SERVER",
		},
		UPDATE_FREEZE_STATUS: JobType{
			value: "UPDATE_FREEZE_STATUS",
		},
		CREATE_SERVER_IMAGE: JobType{
			value: "CREATE_SERVER_IMAGE",
		},
		REINSTALL_OS: JobType{
			value: "REINSTALL_OS",
		},
		CHANGE_SERVER_IMAGE: JobType{
			value: "CHANGE_SERVER_IMAGE",
		},
		REJOIN_DOMAIN: JobType{
			value: "REJOIN_DOMAIN",
		},
		MIGRATE_SERVER: JobType{
			value: "MIGRATE_SERVER",
		},
		UPGRADE_ACCESS_AGENT: JobType{
			value: "UPGRADE_ACCESS_AGENT",
		},
		UPDATE_SERVER_TSVI: JobType{
			value: "UPDATE_SERVER_TSVI",
		},
		SCHEDULED_TASK: JobType{
			value: "SCHEDULED_TASK",
		},
		COLLECT_HDA_LOG: JobType{
			value: "COLLECT_HDA_LOG",
		},
		COLLECT_APS_LOG: JobType{
			value: "COLLECT_APS_LOG",
		},
		CREATE_SERVER_SNAPSHOT: JobType{
			value: "CREATE_SERVER_SNAPSHOT",
		},
		DELETE_SERVER_SNAPSHOT: JobType{
			value: "DELETE_SERVER_SNAPSHOT",
		},
		RESTORE_SERVER_SNAPSHOT: JobType{
			value: "RESTORE_SERVER_SNAPSHOT",
		},
	}
}

func (c JobType) Value() string {
	return c.value
}

func (c JobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobType) UnmarshalJSON(b []byte) error {
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
