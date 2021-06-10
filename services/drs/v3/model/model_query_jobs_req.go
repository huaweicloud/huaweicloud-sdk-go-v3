package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 查询任务列表请求体
type QueryJobsReq struct {
	// 第几页

	CurPage int32 `json:"cur_page"`
	// 每页记录数

	PerPage int32 `json:"per_page"`
	// 迁移场景，migration:实时迁移,sync:实时同步,cloudDataGuard:实时灾备

	DbUseType QueryJobsReqDbUseType `json:"db_use_type"`
	// 引擎类型,mysql：迁移，同步使用。mongodb：迁移使用。cloudDataGuard-mysql：灾备使用。

	EngineType *QueryJobsReqEngineType `json:"engine_type,omitempty"`
	// 企业项目

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// name或id

	Name *string `json:"name,omitempty"`
	// 网络类型

	NetType *QueryJobsReqNetType `json:"net_type,omitempty"`
	// 开启EPS时使用，值为eps

	ServiceName *string `json:"service_name,omitempty"`
	// 状态，CREATING：创建中,CREATE_FAILED: 创建失败,CONFIGURATION: 配置中,STARTJOBING: 启动中,WAITING_FOR_START：等待启动中,START_JOB_FAILED：任务启动失败,FULL_TRANSFER_STARTED：全量迁移中 灾备场景为初始化,FULL_TRANSFER_FAILED：全量迁移失败  灾备场景为初始化失败,FULL_TRANSFER_COMPLETE：全量迁移完成 灾备场景为初始化完成,INCRE_TRANSFER_STARTED：增量迁移中 灾备场景为灾备中,INCRE_TRANSFER_FAILED：增量迁移失败 灾备场景为灾备异常,RELEASE_RESOURCE_STARTED：结束任务中,RELEASE_RESOURCE_FAILED：结束任务失败,RELEASE_RESOURCE_COMPLETE：已结束,CHANGE_JOB_STARTED：任务变更中,CHANGE_JOB_FAILED：任务变更失败,CHILD_TRANSFER_STARTING：子任务启动中,CHILD_TRANSFER_STARTED：子任务迁移中,CHILD_TRANSFER_COMPLETE：子任务迁移完成,CHILD_TRANSFER_FAILED：子任务迁移失败,RELEASE_CHILD_TRANSFER_STARTED：子任务结束中,RELEASE_CHILD_TRANSFER_COMPLETE：子任务已结束

	Status *QueryJobsReqStatus `json:"status,omitempty"`
	// 标签

	Tags map[string]string `json:"tags,omitempty"`
}

func (o QueryJobsReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QueryJobsReq struct{}"
	}

	return strings.Join([]string{"QueryJobsReq", string(data)}, " ")
}

type QueryJobsReqDbUseType struct {
	value string
}

type QueryJobsReqDbUseTypeEnum struct {
	MIGRATION        QueryJobsReqDbUseType
	SYNC             QueryJobsReqDbUseType
	CLOUD_DATA_GUARD QueryJobsReqDbUseType
}

func GetQueryJobsReqDbUseTypeEnum() QueryJobsReqDbUseTypeEnum {
	return QueryJobsReqDbUseTypeEnum{
		MIGRATION: QueryJobsReqDbUseType{
			value: "migration",
		},
		SYNC: QueryJobsReqDbUseType{
			value: "sync",
		},
		CLOUD_DATA_GUARD: QueryJobsReqDbUseType{
			value: "cloudDataGuard",
		},
	}
}

func (c QueryJobsReqDbUseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *QueryJobsReqDbUseType) UnmarshalJSON(b []byte) error {
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

type QueryJobsReqEngineType struct {
	value string
}

type QueryJobsReqEngineTypeEnum struct {
	MYSQL                  QueryJobsReqEngineType
	MONGODB                QueryJobsReqEngineType
	CLOUD_DATA_GUARD_MYSQL QueryJobsReqEngineType
}

func GetQueryJobsReqEngineTypeEnum() QueryJobsReqEngineTypeEnum {
	return QueryJobsReqEngineTypeEnum{
		MYSQL: QueryJobsReqEngineType{
			value: "mysql",
		},
		MONGODB: QueryJobsReqEngineType{
			value: "mongodb",
		},
		CLOUD_DATA_GUARD_MYSQL: QueryJobsReqEngineType{
			value: "cloudDataGuard-mysql",
		},
	}
}

func (c QueryJobsReqEngineType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *QueryJobsReqEngineType) UnmarshalJSON(b []byte) error {
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

type QueryJobsReqNetType struct {
	value string
}

type QueryJobsReqNetTypeEnum struct {
	VPN QueryJobsReqNetType
	VPC QueryJobsReqNetType
	EIP QueryJobsReqNetType
}

func GetQueryJobsReqNetTypeEnum() QueryJobsReqNetTypeEnum {
	return QueryJobsReqNetTypeEnum{
		VPN: QueryJobsReqNetType{
			value: "vpn",
		},
		VPC: QueryJobsReqNetType{
			value: "vpc",
		},
		EIP: QueryJobsReqNetType{
			value: "eip",
		},
	}
}

func (c QueryJobsReqNetType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *QueryJobsReqNetType) UnmarshalJSON(b []byte) error {
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

type QueryJobsReqStatus struct {
	value string
}

type QueryJobsReqStatusEnum struct {
	CREATING                        QueryJobsReqStatus
	CREATE_FAILED                   QueryJobsReqStatus
	CONFIGURATION                   QueryJobsReqStatus
	STARTJOBING                     QueryJobsReqStatus
	WAITING_FOR_START               QueryJobsReqStatus
	START_JOB_FAILED                QueryJobsReqStatus
	FULL_TRANSFER_STARTED           QueryJobsReqStatus
	FULL_TRANSFER_FAILED            QueryJobsReqStatus
	FULL_TRANSFER_COMPLETE          QueryJobsReqStatus
	INCRE_TRANSFER_STARTED          QueryJobsReqStatus
	INCRE_TRANSFER_FAILED           QueryJobsReqStatus
	RELEASE_RESOURCE_STARTED        QueryJobsReqStatus
	RELEASE_RESOURCE_FAILED         QueryJobsReqStatus
	RELEASE_RESOURCE_COMPLETE       QueryJobsReqStatus
	CHANGE_JOB_STARTED              QueryJobsReqStatus
	CHANGE_JOB_FAILED               QueryJobsReqStatus
	CHILD_TRANSFER_STARTING         QueryJobsReqStatus
	CHILD_TRANSFER_STARTED          QueryJobsReqStatus
	CHILD_TRANSFER_COMPLETE         QueryJobsReqStatus
	CHILD_TRANSFER_FAILED           QueryJobsReqStatus
	RELEASE_CHILD_TRANSFER_STARTED  QueryJobsReqStatus
	RELEASE_CHILD_TRANSFER_COMPLETE QueryJobsReqStatus
}

func GetQueryJobsReqStatusEnum() QueryJobsReqStatusEnum {
	return QueryJobsReqStatusEnum{
		CREATING: QueryJobsReqStatus{
			value: "CREATING",
		},
		CREATE_FAILED: QueryJobsReqStatus{
			value: "CREATE_FAILED",
		},
		CONFIGURATION: QueryJobsReqStatus{
			value: "CONFIGURATION",
		},
		STARTJOBING: QueryJobsReqStatus{
			value: "STARTJOBING",
		},
		WAITING_FOR_START: QueryJobsReqStatus{
			value: "WAITING_FOR_START",
		},
		START_JOB_FAILED: QueryJobsReqStatus{
			value: "START_JOB_FAILED",
		},
		FULL_TRANSFER_STARTED: QueryJobsReqStatus{
			value: "FULL_TRANSFER_STARTED",
		},
		FULL_TRANSFER_FAILED: QueryJobsReqStatus{
			value: "FULL_TRANSFER_FAILED",
		},
		FULL_TRANSFER_COMPLETE: QueryJobsReqStatus{
			value: "FULL_TRANSFER_COMPLETE",
		},
		INCRE_TRANSFER_STARTED: QueryJobsReqStatus{
			value: "INCRE_TRANSFER_STARTED",
		},
		INCRE_TRANSFER_FAILED: QueryJobsReqStatus{
			value: "INCRE_TRANSFER_FAILED",
		},
		RELEASE_RESOURCE_STARTED: QueryJobsReqStatus{
			value: "RELEASE_RESOURCE_STARTED",
		},
		RELEASE_RESOURCE_FAILED: QueryJobsReqStatus{
			value: "RELEASE_RESOURCE_FAILED",
		},
		RELEASE_RESOURCE_COMPLETE: QueryJobsReqStatus{
			value: "RELEASE_RESOURCE_COMPLETE",
		},
		CHANGE_JOB_STARTED: QueryJobsReqStatus{
			value: "CHANGE_JOB_STARTED",
		},
		CHANGE_JOB_FAILED: QueryJobsReqStatus{
			value: "CHANGE_JOB_FAILED",
		},
		CHILD_TRANSFER_STARTING: QueryJobsReqStatus{
			value: "CHILD_TRANSFER_STARTING",
		},
		CHILD_TRANSFER_STARTED: QueryJobsReqStatus{
			value: "CHILD_TRANSFER_STARTED",
		},
		CHILD_TRANSFER_COMPLETE: QueryJobsReqStatus{
			value: "CHILD_TRANSFER_COMPLETE",
		},
		CHILD_TRANSFER_FAILED: QueryJobsReqStatus{
			value: "CHILD_TRANSFER_FAILED",
		},
		RELEASE_CHILD_TRANSFER_STARTED: QueryJobsReqStatus{
			value: "RELEASE_CHILD_TRANSFER_STARTED",
		},
		RELEASE_CHILD_TRANSFER_COMPLETE: QueryJobsReqStatus{
			value: "RELEASE_CHILD_TRANSFER_COMPLETE",
		},
	}
}

func (c QueryJobsReqStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *QueryJobsReqStatus) UnmarshalJSON(b []byte) error {
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
