package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DbOffsiteBackupPolicyBody struct {

	// 指定备份的类型。取值如下：auto：自动全量备份。incremental：自动增量备份。all：同时设置所有备份类型。Db: 数据库。Log：日志
	BackupType DbOffsiteBackupPolicyBodyBackupType `json:"backup_type"`

	// 指定已生成的备份文件可以保存的天数。取值范围：0～1825。保存天数设置为0时，表示关闭跨区域备份策略。
	KeepDays int32 `json:"keep_days"`

	// 设置跨区域备份策略的目标区域ID。
	DestinationRegion *string `json:"destination_region,omitempty"`

	// 设置跨区域备份策略的目标project ID。
	DestinationProjectId *string `json:"destination_project_id,omitempty"`
}

func (o DbOffsiteBackupPolicyBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DbOffsiteBackupPolicyBody struct{}"
	}

	return strings.Join([]string{"DbOffsiteBackupPolicyBody", string(data)}, " ")
}

type DbOffsiteBackupPolicyBodyBackupType struct {
	value string
}

type DbOffsiteBackupPolicyBodyBackupTypeEnum struct {
	AUTO        DbOffsiteBackupPolicyBodyBackupType
	INCREMENTAL DbOffsiteBackupPolicyBodyBackupType
	ALL         DbOffsiteBackupPolicyBodyBackupType
	DB          DbOffsiteBackupPolicyBodyBackupType
	LOG         DbOffsiteBackupPolicyBodyBackupType
}

func GetDbOffsiteBackupPolicyBodyBackupTypeEnum() DbOffsiteBackupPolicyBodyBackupTypeEnum {
	return DbOffsiteBackupPolicyBodyBackupTypeEnum{
		AUTO: DbOffsiteBackupPolicyBodyBackupType{
			value: "auto",
		},
		INCREMENTAL: DbOffsiteBackupPolicyBodyBackupType{
			value: "incremental",
		},
		ALL: DbOffsiteBackupPolicyBodyBackupType{
			value: "all",
		},
		DB: DbOffsiteBackupPolicyBodyBackupType{
			value: "Db",
		},
		LOG: DbOffsiteBackupPolicyBodyBackupType{
			value: "Log",
		},
	}
}

func (c DbOffsiteBackupPolicyBodyBackupType) Value() string {
	return c.value
}

func (c DbOffsiteBackupPolicyBodyBackupType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DbOffsiteBackupPolicyBodyBackupType) UnmarshalJSON(b []byte) error {
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
