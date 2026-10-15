package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobBaseInfo 创建任务基本信息体。
type JobBaseInfo struct {

	// **参数解释**：  任务名称，用于标识和区分不同的数据复制任务。  **约束限制**：  - 长度在4到50位之间，不区分大小写。 - 可以包含字母、数字、中划线或下划线，不能包括其他特殊字符。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**：  DRS任务场景类型，用于区分不同的数据复制场景。  **约束限制**：  不涉及。  **取值范围**：  - migration：实时迁移 - sync：实时同步  **默认取值**：  不涉及
	JobType *JobBaseInfoJobType `json:"job_type,omitempty"`

	// **参数解释**：  灾备类型是否双主灾备，指定灾备场景下是否启用双向写入功能。  **约束限制**：  - job_type为cloudDataGuard时必填。  **取值范围**：  - true：双主灾备 - false：非双主灾备  **默认取值**：  false
	MultiWrite *bool `json:"multi_write,omitempty"`

	// **参数解释**：  DRS任务引擎类型。  **约束限制**：  不涉及。  **取值范围**：  - oracle-to-gaussdbv5：Oracle同步到GaussDB分布式版，实时同步场景使用 - redis-to-gaussredis：Redis同步到GeminiDB Redis，实时迁移场景使用 - rediscluster-to-gaussredis：Redis集群同步到GeminiDB Redis，实时迁移场景使用 - mysql-to-mysql：MySQL同步到MySQL，实时同步场景使用  **默认取值**：  不涉及
	EngineType *JobBaseInfoEngineType `json:"engine_type,omitempty"`

	// **参数解释**：  任务的迁移方向。  **约束限制**：  不涉及。  **取值范围**：  - up：入云，灾备场景时对应本云为备 - down：出云，灾备场景时对应本云为主 - non-dbs：自建  **默认取值**：  不涉及
	JobDirection *JobBaseInfoJobDirection `json:"job_direction,omitempty"`

	// **参数解释**：  任务迁移模式，指定数据复制任务的执行方式。  **约束限制**：  不涉及。  **取值范围**：  - FULL_TRANS：全量迁移，仅迁移历史数据 - FULL_INCR_TRANS：全量+增量迁移，先迁移历史数据再持续同步增量 - INCR_TRANS：增量迁移，仅同步数据变更  **默认取值**：  不涉及
	TaskType *JobBaseInfoTaskType `json:"task_type,omitempty"`

	// **参数解释**：  任务使用的网络类型。  **约束限制**：  - 灾备场景不支持选择VPC网络。  **取值范围**：  - eip：公网网络 - vpc：VPC网络 - vpn：VPN、专线网络  **默认取值**：  不涉及
	NetType *JobBaseInfoNetType `json:"net_type,omitempty"`

	// **参数解释**：  任务的计费模式，指定数据复制任务的付费方式。  **约束限制**：  不涉及。  **取值范围**：  - period：包周期 - on_demand：按需  **默认取值**：  on_demand
	ChargingMode *JobBaseInfoChargingMode `json:"charging_mode,omitempty"`

	// **参数解释**：  企业项目ID，用于按企业项目过滤查询结果。  **约束限制**：  - 不传值时设置为null，不能设置为空字符串。 - 当企业项目开关打开时，可以传值。  **取值范围**：  不涉及。  **默认取值**：  \"0\"，表示default企业项目
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**：  任务描述，用于对数据复制任务添加补充说明信息。  **约束限制**：  - 长度不能超过256位。 - 不能包含!<>&'\"\\特殊字符。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**：  任务定时启动时间，设置后任务将在指定时间自动启动。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	StartTime *string `json:"start_time,omitempty"`

	// 任务处于异常状态一段时间后，将会自动结束。单位为天。(范围14-100)，不传默认为14天。
	ExpiredDays *string `json:"expired_days,omitempty"`

	// **参数解释**：  标签信息，用于为数据复制任务添加标签以便分类管理。  **约束限制**：  - 最多添加10个标签。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Tags *[]ResourceTag `json:"tags,omitempty"`

	// **参数解释**：  是否开启云数据库RDS for MySQL/MariaDB的binlog快速清理，开启后可自动清理已同步的binlog以释放磁盘空间。  **约束限制**：  - 仅MySQL/MariaDB引擎类型适用。  **取值范围**：  - true：开启binlog快速清理 - false：不开启binlog快速清理  **默认取值**：  false
	IsOpenFastClean *bool `json:"is_open_fast_clean,omitempty"`
}

func (o JobBaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobBaseInfo struct{}"
	}

	return strings.Join([]string{"JobBaseInfo", string(data)}, " ")
}

type JobBaseInfoJobType struct {
	value string
}

type JobBaseInfoJobTypeEnum struct {
	MIGRATION        JobBaseInfoJobType
	SYNC             JobBaseInfoJobType
	CLOUD_DATA_GUARD JobBaseInfoJobType
}

func GetJobBaseInfoJobTypeEnum() JobBaseInfoJobTypeEnum {
	return JobBaseInfoJobTypeEnum{
		MIGRATION: JobBaseInfoJobType{
			value: "migration",
		},
		SYNC: JobBaseInfoJobType{
			value: "sync",
		},
		CLOUD_DATA_GUARD: JobBaseInfoJobType{
			value: "cloudDataGuard",
		},
	}
}

func (c JobBaseInfoJobType) Value() string {
	return c.value
}

func (c JobBaseInfoJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobBaseInfoJobType) UnmarshalJSON(b []byte) error {
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

type JobBaseInfoEngineType struct {
	value string
}

type JobBaseInfoEngineTypeEnum struct {
	ORACLE_TO_GAUSSDBV5        JobBaseInfoEngineType
	REDIS_TO_GAUSSREDIS        JobBaseInfoEngineType
	REDISCLUSTER_TO_GAUSSREDIS JobBaseInfoEngineType
	MYSQL_TO_MYSQL             JobBaseInfoEngineType
}

func GetJobBaseInfoEngineTypeEnum() JobBaseInfoEngineTypeEnum {
	return JobBaseInfoEngineTypeEnum{
		ORACLE_TO_GAUSSDBV5: JobBaseInfoEngineType{
			value: "oracle-to-gaussdbv5",
		},
		REDIS_TO_GAUSSREDIS: JobBaseInfoEngineType{
			value: "redis-to-gaussredis",
		},
		REDISCLUSTER_TO_GAUSSREDIS: JobBaseInfoEngineType{
			value: "rediscluster-to-gaussredis",
		},
		MYSQL_TO_MYSQL: JobBaseInfoEngineType{
			value: "mysql-to-mysql",
		},
	}
}

func (c JobBaseInfoEngineType) Value() string {
	return c.value
}

func (c JobBaseInfoEngineType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobBaseInfoEngineType) UnmarshalJSON(b []byte) error {
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

type JobBaseInfoJobDirection struct {
	value string
}

type JobBaseInfoJobDirectionEnum struct {
	UP      JobBaseInfoJobDirection
	DOWN    JobBaseInfoJobDirection
	NON_DBS JobBaseInfoJobDirection
}

func GetJobBaseInfoJobDirectionEnum() JobBaseInfoJobDirectionEnum {
	return JobBaseInfoJobDirectionEnum{
		UP: JobBaseInfoJobDirection{
			value: "up",
		},
		DOWN: JobBaseInfoJobDirection{
			value: "down",
		},
		NON_DBS: JobBaseInfoJobDirection{
			value: "non-dbs",
		},
	}
}

func (c JobBaseInfoJobDirection) Value() string {
	return c.value
}

func (c JobBaseInfoJobDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobBaseInfoJobDirection) UnmarshalJSON(b []byte) error {
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

type JobBaseInfoTaskType struct {
	value string
}

type JobBaseInfoTaskTypeEnum struct {
	FULL_TRANS      JobBaseInfoTaskType
	FULL_INCR_TRANS JobBaseInfoTaskType
	INCR_TRANS      JobBaseInfoTaskType
}

func GetJobBaseInfoTaskTypeEnum() JobBaseInfoTaskTypeEnum {
	return JobBaseInfoTaskTypeEnum{
		FULL_TRANS: JobBaseInfoTaskType{
			value: "FULL_TRANS",
		},
		FULL_INCR_TRANS: JobBaseInfoTaskType{
			value: "FULL_INCR_TRANS",
		},
		INCR_TRANS: JobBaseInfoTaskType{
			value: "INCR_TRANS",
		},
	}
}

func (c JobBaseInfoTaskType) Value() string {
	return c.value
}

func (c JobBaseInfoTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobBaseInfoTaskType) UnmarshalJSON(b []byte) error {
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

type JobBaseInfoNetType struct {
	value string
}

type JobBaseInfoNetTypeEnum struct {
	EIP JobBaseInfoNetType
	VPC JobBaseInfoNetType
	VPN JobBaseInfoNetType
}

func GetJobBaseInfoNetTypeEnum() JobBaseInfoNetTypeEnum {
	return JobBaseInfoNetTypeEnum{
		EIP: JobBaseInfoNetType{
			value: "eip",
		},
		VPC: JobBaseInfoNetType{
			value: "vpc",
		},
		VPN: JobBaseInfoNetType{
			value: "vpn",
		},
	}
}

func (c JobBaseInfoNetType) Value() string {
	return c.value
}

func (c JobBaseInfoNetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobBaseInfoNetType) UnmarshalJSON(b []byte) error {
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

type JobBaseInfoChargingMode struct {
	value string
}

type JobBaseInfoChargingModeEnum struct {
	PERIOD    JobBaseInfoChargingMode
	ON_DEMAND JobBaseInfoChargingMode
}

func GetJobBaseInfoChargingModeEnum() JobBaseInfoChargingModeEnum {
	return JobBaseInfoChargingModeEnum{
		PERIOD: JobBaseInfoChargingMode{
			value: "period",
		},
		ON_DEMAND: JobBaseInfoChargingMode{
			value: "on_demand",
		},
	}
}

func (c JobBaseInfoChargingMode) Value() string {
	return c.value
}

func (c JobBaseInfoChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobBaseInfoChargingMode) UnmarshalJSON(b []byte) error {
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
