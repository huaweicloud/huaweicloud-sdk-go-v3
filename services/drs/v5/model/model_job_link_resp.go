package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 可用链路返回体。
type JobLinkResp struct {

	// 任务场景。取值： - migration：实时迁移。 - sync：实时同步。 - cloudDataGuard：实时灾备。
	JobType JobLinkRespJobType `json:"job_type"`

	// 引擎类型。取值： - oracle-to-gaussdbv5：Oracle同步到GaussDB分布式版，实时同步场景使用。
	EngineType JobLinkRespEngineType `json:"engine_type"`

	// 源数据库实例类型。取值： - offline：自建数据库。 - ecs：华为云ECS自建数据库。 - cloud：华为云数据库。
	SourceEndpointType JobLinkRespSourceEndpointType `json:"source_endpoint_type"`

	// 目标数据库实例类型。取值： - offline：自建数据库。 - ecs：华为云ECS自建数据库。 - cloud：华为云数据库。
	TargetEndpointType JobLinkRespTargetEndpointType `json:"target_endpoint_type"`

	// 迁移方向。取值： - up：入云 ，灾备场景时对应本云为备。 - down：出云，灾备场景时对应本云为主。 - non-dbs：自建。
	JobDirection JobLinkRespJobDirection `json:"job_direction"`

	// 网络类型。取值： - eip：公网网络。 - vpc：VPC网络，灾备场景不支持选择VPC网络。 - vpn：VPN、专线网络。
	NetType JobLinkRespNetType `json:"net_type"`

	// 迁移模式。取值： - FULL_TRANS ：全量。 - FULL_INCR_TRANS：全量+增量。 - INCR_TRANS：增量。
	TaskTypes []JobLinkRespTaskTypes `json:"task_types"`

	// 引擎实例模式。取值： - Single：单机模式。 - Ha：主备模式。 - Cluster：集群模式。 - Sharding：分片模式。 - Independent：GaussDB独立部署模式。
	ClusterModes []JobLinkRespClusterModes `json:"cluster_modes"`

	// 链路描述。
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

type JobLinkRespEngineType struct {
	value string
}

type JobLinkRespEngineTypeEnum struct {
	ORACLE_TO_GAUSSDBV5 JobLinkRespEngineType
}

func GetJobLinkRespEngineTypeEnum() JobLinkRespEngineTypeEnum {
	return JobLinkRespEngineTypeEnum{
		ORACLE_TO_GAUSSDBV5: JobLinkRespEngineType{
			value: "oracle-to-gaussdbv5",
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
