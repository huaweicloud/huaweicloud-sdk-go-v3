package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobBaseInfo 创建任务基本信息体。
type JobBaseInfo struct {

	// 任务名称。 约束：任务名称在4位到50位之间，不区分大小写，可以包含字母、数字、中划线或下划线，不能包括其他特殊字符。 - 最小长度：4 - 最大长度：50
	Name *string `json:"name,omitempty"`

	// 任务场景。取值： - migration：实时迁移。 - sync：实时同步。 - cloudDataGuard：实时灾备。
	JobType *JobBaseInfoJobType `json:"job_type,omitempty"`

	// 灾备类型是否双主灾备。说明： - job_type 是cloudDataGuard时，必填，灾备类型是双主灾备时，multi_write取值true, 否则为false。 - job_type 是其他类型时，multi_write是非必选参数。
	MultiWrite *bool `json:"multi_write,omitempty"`

	// 引擎类型。取值： - oracle-to-gaussdbv5：Oracle同步到GaussDB分布式版，实时同步场景使用。 - redis-to-gaussredis：Redis同步到GeminiDB Redis，实时迁移场景使用。 - rediscluster-to-gaussredis：Redis集群同步到GeminiDB Redis，实时迁移场景使用。 - mysql-to-mysql：MySQL同步到MySQL，实时同步场景使用。
	EngineType *JobBaseInfoEngineType `json:"engine_type,omitempty"`

	// 迁移方向。取值： - up：入云 ，灾备场景时对应本云为备。 - down：出云，灾备场景时对应本云为主。 - non-dbs：自建。
	JobDirection *JobBaseInfoJobDirection `json:"job_direction,omitempty"`

	// 迁移模式。取值： - FULL_TRANS ：全量。 - FULL_INCR_TRANS：全量+增量。 - INCR_TRANS：增量。
	TaskType *JobBaseInfoTaskType `json:"task_type,omitempty"`

	// 网络类型。取值： - eip：公网网络。 - vpc：VPC网络，灾备场景不支持选择VPC网络。 - vpn：VPN、专线网络。
	NetType *JobBaseInfoNetType `json:"net_type,omitempty"`

	// 计费模式，默认按需。取值： - period：包周期。 - on_demand：按需。
	ChargingMode *JobBaseInfoChargingMode `json:"charging_mode,omitempty"`

	// 企业项目ID。 缺省值：\"0\"，表示\"default\"企业项目。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 任务描述。 约束：任务描述不能超过256位，且不能包含!<>&'\"\\特殊字符。
	Description *string `json:"description,omitempty"`

	// 任务定时启动时间。
	StartTime *string `json:"start_time,omitempty"`

	// 任务处于异常状态一段时间后，将会自动结束。单位为天。(范围14-100)，不传默认为14天。
	ExpiredDays *string `json:"expired_days,omitempty"`

	// 标签信息，最多添加10个标签。
	Tags *[]ResourceTag `json:"tags,omitempty"`

	// 是否开启云数据库RDS for MySQL/MariaDB的binlog快速清理。
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
