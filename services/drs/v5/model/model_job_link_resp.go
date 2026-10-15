package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobLinkResp 可用链路返回体。
type JobLinkResp struct {

	// **参数解释**：  DRS任务场景类型，用于区分不同的数据复制场景。  **约束限制**：  不涉及。  **取值范围**：  - migration：实时迁移 - sync：实时同步 - cloudDataGuard：实时灾备  **默认取值**：  不涉及
	JobType JobLinkRespJobType `json:"job_type"`

	// **参数解释**：  DRS任务引擎类型，标识数据复制任务所使用的数据库引擎组合。  **约束限制**：  不涉及。  **取值范围**：  - oracle-to-gaussdbv5：Oracle同步到GaussDB分布式版，实时同步场景使用 - redis-to-gaussredis：Redis迁移到GeminiDB Redis，实时迁移场景使用 - rediscluster-to-gaussredis：Redis集群迁移到GeminiDB Redis，实时迁移场景使用  **默认取值**：  不涉及
	EngineType JobLinkRespEngineType `json:"engine_type"`

	// **参数解释**：  源数据库实例类型，标识源数据库的部署方式。  **约束限制**：  不涉及。  **取值范围**：  - offline：自建数据库 - ecs：华为云ECS自建数据库 - cloud：华为云数据库  **默认取值**：  不涉及
	SourceEndpointType JobLinkRespSourceEndpointType `json:"source_endpoint_type"`

	// **参数解释**：  目标数据库实例类型，标识目标数据库的部署方式。  **约束限制**：  不涉及。  **取值范围**：  - offline：自建数据库 - ecs：华为云ECS自建数据库 - cloud：华为云数据库  **默认取值**：  不涉及
	TargetEndpointType JobLinkRespTargetEndpointType `json:"target_endpoint_type"`

	// **参数解释**：  任务的迁移方向。  **约束限制**：  不涉及。  **取值范围**：  - up：入云，灾备场景时对应本云为备 - down：出云，灾备场景时对应本云为主 - non-dbs：自建  **默认取值**：  不涉及
	JobDirection JobLinkRespJobDirection `json:"job_direction"`

	// **参数解释**：  任务使用的网络类型。  **约束限制**：  不涉及。  **取值范围**：  - eip：公网网络 - vpc：VPC网络，灾备场景不支持选择VPC网络 - vpn：VPN、专线网络  **默认取值**：  不涉及
	NetType JobLinkRespNetType `json:"net_type"`

	// **参数解释**：  可用链路支持的迁移模式列表，表示该链路支持的数据迁移执行方式。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	TaskTypes []JobLinkRespTaskTypes `json:"task_types"`

	// **参数解释**：  可用链路支持的引擎实例模式列表，表示该链路支持的数据库部署架构类型。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	ClusterModes []JobLinkRespClusterModes `json:"cluster_modes"`

	// **参数解释**：  链路的描述信息，用于说明该链路的适用场景或功能特点。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Description string `json:"description"`
}

func (o JobLinkResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobLinkResp struct{}"
	}

	return strings.Join([]string{"JobLinkResp", string(data)}, " ")
}

type JobLinkRespJobType struct {
	value string
}

type JobLinkRespJobTypeEnum struct {
	MIGRATION        JobLinkRespJobType
	SYNC             JobLinkRespJobType
	CLOUD_DATA_GUARD JobLinkRespJobType
}

func GetJobLinkRespJobTypeEnum() JobLinkRespJobTypeEnum {
	return JobLinkRespJobTypeEnum{
		MIGRATION: JobLinkRespJobType{
			value: "migration",
		},
		SYNC: JobLinkRespJobType{
			value: "sync",
		},
		CLOUD_DATA_GUARD: JobLinkRespJobType{
			value: "cloudDataGuard",
		},
	}
}

func (c JobLinkRespJobType) Value() string {
	return c.value
}

func (c JobLinkRespJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobLinkRespJobType) UnmarshalJSON(b []byte) error {
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

type JobLinkRespEngineType struct {
	value string
}

type JobLinkRespEngineTypeEnum struct {
	ORACLE_TO_GAUSSDBV5        JobLinkRespEngineType
	REDIS_TO_GAUSSREDIS        JobLinkRespEngineType
	REDISCLUSTER_TO_GAUSSREDIS JobLinkRespEngineType
}

func GetJobLinkRespEngineTypeEnum() JobLinkRespEngineTypeEnum {
	return JobLinkRespEngineTypeEnum{
		ORACLE_TO_GAUSSDBV5: JobLinkRespEngineType{
			value: "oracle-to-gaussdbv5",
		},
		REDIS_TO_GAUSSREDIS: JobLinkRespEngineType{
			value: "redis-to-gaussredis",
		},
		REDISCLUSTER_TO_GAUSSREDIS: JobLinkRespEngineType{
			value: "rediscluster-to-gaussredis",
		},
	}
}

func (c JobLinkRespEngineType) Value() string {
	return c.value
}

func (c JobLinkRespEngineType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobLinkRespEngineType) UnmarshalJSON(b []byte) error {
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

type JobLinkRespSourceEndpointType struct {
	value string
}

type JobLinkRespSourceEndpointTypeEnum struct {
	OFFLINE JobLinkRespSourceEndpointType
	ECS     JobLinkRespSourceEndpointType
	CLOUD   JobLinkRespSourceEndpointType
}

func GetJobLinkRespSourceEndpointTypeEnum() JobLinkRespSourceEndpointTypeEnum {
	return JobLinkRespSourceEndpointTypeEnum{
		OFFLINE: JobLinkRespSourceEndpointType{
			value: "offline",
		},
		ECS: JobLinkRespSourceEndpointType{
			value: "ecs",
		},
		CLOUD: JobLinkRespSourceEndpointType{
			value: "cloud",
		},
	}
}

func (c JobLinkRespSourceEndpointType) Value() string {
	return c.value
}

func (c JobLinkRespSourceEndpointType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobLinkRespSourceEndpointType) UnmarshalJSON(b []byte) error {
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

type JobLinkRespTargetEndpointType struct {
	value string
}

type JobLinkRespTargetEndpointTypeEnum struct {
	OFFLINE JobLinkRespTargetEndpointType
	ECS     JobLinkRespTargetEndpointType
	CLOUD   JobLinkRespTargetEndpointType
}

func GetJobLinkRespTargetEndpointTypeEnum() JobLinkRespTargetEndpointTypeEnum {
	return JobLinkRespTargetEndpointTypeEnum{
		OFFLINE: JobLinkRespTargetEndpointType{
			value: "offline",
		},
		ECS: JobLinkRespTargetEndpointType{
			value: "ecs",
		},
		CLOUD: JobLinkRespTargetEndpointType{
			value: "cloud",
		},
	}
}

func (c JobLinkRespTargetEndpointType) Value() string {
	return c.value
}

func (c JobLinkRespTargetEndpointType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobLinkRespTargetEndpointType) UnmarshalJSON(b []byte) error {
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

type JobLinkRespJobDirection struct {
	value string
}

type JobLinkRespJobDirectionEnum struct {
	UP      JobLinkRespJobDirection
	DOWN    JobLinkRespJobDirection
	NON_DBS JobLinkRespJobDirection
}

func GetJobLinkRespJobDirectionEnum() JobLinkRespJobDirectionEnum {
	return JobLinkRespJobDirectionEnum{
		UP: JobLinkRespJobDirection{
			value: "up",
		},
		DOWN: JobLinkRespJobDirection{
			value: "down",
		},
		NON_DBS: JobLinkRespJobDirection{
			value: "non-dbs",
		},
	}
}

func (c JobLinkRespJobDirection) Value() string {
	return c.value
}

func (c JobLinkRespJobDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobLinkRespJobDirection) UnmarshalJSON(b []byte) error {
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

type JobLinkRespNetType struct {
	value string
}

type JobLinkRespNetTypeEnum struct {
	EIP JobLinkRespNetType
	VPC JobLinkRespNetType
	VPN JobLinkRespNetType
}

func GetJobLinkRespNetTypeEnum() JobLinkRespNetTypeEnum {
	return JobLinkRespNetTypeEnum{
		EIP: JobLinkRespNetType{
			value: "eip",
		},
		VPC: JobLinkRespNetType{
			value: "vpc",
		},
		VPN: JobLinkRespNetType{
			value: "vpn",
		},
	}
}

func (c JobLinkRespNetType) Value() string {
	return c.value
}

func (c JobLinkRespNetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobLinkRespNetType) UnmarshalJSON(b []byte) error {
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

type JobLinkRespTaskTypes struct {
	value string
}

type JobLinkRespTaskTypesEnum struct {
	FULL_TRANS      JobLinkRespTaskTypes
	FULL_INCR_TRANS JobLinkRespTaskTypes
	INCR_TRANS      JobLinkRespTaskTypes
}

func GetJobLinkRespTaskTypesEnum() JobLinkRespTaskTypesEnum {
	return JobLinkRespTaskTypesEnum{
		FULL_TRANS: JobLinkRespTaskTypes{
			value: "FULL_TRANS",
		},
		FULL_INCR_TRANS: JobLinkRespTaskTypes{
			value: "FULL_INCR_TRANS",
		},
		INCR_TRANS: JobLinkRespTaskTypes{
			value: "INCR_TRANS",
		},
	}
}

func (c JobLinkRespTaskTypes) Value() string {
	return c.value
}

func (c JobLinkRespTaskTypes) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobLinkRespTaskTypes) UnmarshalJSON(b []byte) error {
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

type JobLinkRespClusterModes struct {
	value string
}

type JobLinkRespClusterModesEnum struct {
	SINGLE      JobLinkRespClusterModes
	HA          JobLinkRespClusterModes
	CLUSTER     JobLinkRespClusterModes
	SHARDING    JobLinkRespClusterModes
	INDEPENDENT JobLinkRespClusterModes
}

func GetJobLinkRespClusterModesEnum() JobLinkRespClusterModesEnum {
	return JobLinkRespClusterModesEnum{
		SINGLE: JobLinkRespClusterModes{
			value: "Single",
		},
		HA: JobLinkRespClusterModes{
			value: "Ha",
		},
		CLUSTER: JobLinkRespClusterModes{
			value: "Cluster",
		},
		SHARDING: JobLinkRespClusterModes{
			value: "Sharding",
		},
		INDEPENDENT: JobLinkRespClusterModes{
			value: "Independent",
		},
	}
}

func (c JobLinkRespClusterModes) Value() string {
	return c.value
}

func (c JobLinkRespClusterModes) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobLinkRespClusterModes) UnmarshalJSON(b []byte) error {
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
