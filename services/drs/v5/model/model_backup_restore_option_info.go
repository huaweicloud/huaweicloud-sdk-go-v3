package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BackupRestoreOptionInfo 备份迁移参数。
type BackupRestoreOptionInfo struct {

	// 是否覆盖目标数据库，不填默认为false。  值为“true”表示覆盖。  值为“false”表示不覆盖。
	IsCover *bool `json:"is_cover,omitempty"`

	// 是否默认恢复，默认恢复所有数据库。  “true”是代表默认恢复还原备份文件中的全部数据库。  “false”代表需要在db_names这个字段中指定需要恢复的数据库名。
	IsDefaultRestore *bool `json:"is_default_restore,omitempty"`

	// 一次典型的增量恢复过程，会涉及多次恢复增量备份。每个增量备份恢复均会使目标数据库保持还原中状态，此时数据库不可读写，直至最后一个增量备份恢复完成后，数据库才能变成可用状态。此后数据库将无法继续进行增量恢复，所以确定为最后一个备份的场景有：  一次性全量迁移，后续将不再进行增量恢复时，设置该字段值为“true”。  增量恢复流程中，确定为最后割接阶段的最后一个增量备份时，设置该字段值为“false”。
	IsLastBackup bool `json:"is_last_backup"`

	// 是否执行预校验。 true：执行。 false：不执行。
	IsPrecheck bool `json:"is_precheck"`

	// 恢复模式：  “full”表示全量迁移。  “incre”表示增量迁移 。
	RecoveryMode BackupRestoreOptionInfoRecoveryMode `json:"recovery_mode"`

	// 数据库名称。
	DbNames *[]string `json:"db_names,omitempty"`

	// 该字段为一个map，目前使用格式key是\"\"，value是新数据库名字。 该功能将忽略备份文件中原有的数据库名，通过DRS将其恢复为指定的新数据库名。 使用条件： - 备份文件中只有一个数据库。 - 备份文件是全量备份类型（待恢复备份类型选择：全量备份），且是一次性恢复（最后一个备份选择：是）。
	ResetDbNameMap map[string]string `json:"reset_db_name_map,omitempty"`

	// 该参数控制使用OBS桶中备份文件恢复时，当任务结束时是否删除下载到RDS for SQL server磁盘中的OBS备份文件，默认删除。 - true 删除 - false 不删除
	IsDeleteBackupFile *bool `json:"is_delete_backup_file,omitempty"`
}

func (o BackupRestoreOptionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupRestoreOptionInfo struct{}"
	}

	return strings.Join([]string{"BackupRestoreOptionInfo", string(data)}, " ")
}

type BackupRestoreOptionInfoRecoveryMode struct {
	value string
}

type BackupRestoreOptionInfoRecoveryModeEnum struct {
	FULL  BackupRestoreOptionInfoRecoveryMode
	INCRE BackupRestoreOptionInfoRecoveryMode
}

func GetBackupRestoreOptionInfoRecoveryModeEnum() BackupRestoreOptionInfoRecoveryModeEnum {
	return BackupRestoreOptionInfoRecoveryModeEnum{
		FULL: BackupRestoreOptionInfoRecoveryMode{
			value: "full",
		},
		INCRE: BackupRestoreOptionInfoRecoveryMode{
			value: "incre",
		},
	}
}

func (c BackupRestoreOptionInfoRecoveryMode) Value() string {
	return c.value
}

func (c BackupRestoreOptionInfoRecoveryMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackupRestoreOptionInfoRecoveryMode) UnmarshalJSON(b []byte) error {
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
