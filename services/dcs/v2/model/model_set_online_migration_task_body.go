package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 配置在线数据迁移任务结构体
type SetOnlineMigrationTaskBody struct {

	// 迁移方式，包括全量迁移和增量迁移两种类型。 - 全量迁移：该模式为Redis的一次性迁移，适用于可中断业务的迁移场景。   全量迁移过程中，如果源Redis有数据更新，这部分更新数据不会被迁移到目标Redis。 - 增量迁移：该模式为Redis的持续性迁移，适用于对业务中断敏感的迁移场景。   增量迁移阶段通过解析日志等技术， 持续保持源Redis和目标端Redis的数据一致。 取值范围： - full_amount_migration：表示全量迁移。 - incremental_migration：表示增量迁移。
	MigrationMethod SetOnlineMigrationTaskBodyMigrationMethod `json:"migration_method"`

	// 自动重连，根据参数决定是否自动重连。 自动重连模式在遇到网络等异常情况时，会无限自动重试。 自动重连模式在无法进行增量同步时，会触发全量同步，增加带宽占用，请谨慎选择。 取值范围： - auto：自动重连。 - manual：手动重连。
	ResumeMode SetOnlineMigrationTaskBodyResumeMode `json:"resume_mode"`

	// 带宽限制，当迁移方式为增量迁移时，为保证业务正常运行，您可以启用带宽限制功能，当数据同步速度达到带宽限制时，将限制同步速度的继续增长。 -限制为MB/s -取值范围：1-10,241(大于0小于10,241的整数)
	BandwidthLimitMb *string `json:"bandwidth_limit_mb,omitempty"`

	SourceInstance *ConfigMigrationInstanceBody `json:"source_instance"`

	TargetInstance *ConfigMigrationInstanceBody `json:"target_instance"`
}

func (o SetOnlineMigrationTaskBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetOnlineMigrationTaskBody struct{}"
	}

	return strings.Join([]string{"SetOnlineMigrationTaskBody", string(data)}, " ")
}

type SetOnlineMigrationTaskBodyMigrationMethod struct {
	value string
}

type SetOnlineMigrationTaskBodyMigrationMethodEnum struct {
	FULL_AMOUNT_MIGRATION SetOnlineMigrationTaskBodyMigrationMethod
	INCREMENTAL_MIGRATION SetOnlineMigrationTaskBodyMigrationMethod
}

func GetSetOnlineMigrationTaskBodyMigrationMethodEnum() SetOnlineMigrationTaskBodyMigrationMethodEnum {
	return SetOnlineMigrationTaskBodyMigrationMethodEnum{
		FULL_AMOUNT_MIGRATION: SetOnlineMigrationTaskBodyMigrationMethod{
			value: "full_amount_migration",
		},
		INCREMENTAL_MIGRATION: SetOnlineMigrationTaskBodyMigrationMethod{
			value: "incremental_migration",
		},
	}
}

func (c SetOnlineMigrationTaskBodyMigrationMethod) Value() string {
	return c.value
}

func (c SetOnlineMigrationTaskBodyMigrationMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SetOnlineMigrationTaskBodyMigrationMethod) UnmarshalJSON(b []byte) error {
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

type SetOnlineMigrationTaskBodyResumeMode struct {
	value string
}

type SetOnlineMigrationTaskBodyResumeModeEnum struct {
	AUTO   SetOnlineMigrationTaskBodyResumeMode
	MANUAL SetOnlineMigrationTaskBodyResumeMode
}

func GetSetOnlineMigrationTaskBodyResumeModeEnum() SetOnlineMigrationTaskBodyResumeModeEnum {
	return SetOnlineMigrationTaskBodyResumeModeEnum{
		AUTO: SetOnlineMigrationTaskBodyResumeMode{
			value: "auto",
		},
		MANUAL: SetOnlineMigrationTaskBodyResumeMode{
			value: "manual",
		},
	}
}

func (c SetOnlineMigrationTaskBodyResumeMode) Value() string {
	return c.value
}

func (c SetOnlineMigrationTaskBodyResumeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SetOnlineMigrationTaskBodyResumeMode) UnmarshalJSON(b []byte) error {
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
