package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobInfoRequest 作业详情请求体
type JobInfoRequest struct {

	// 作业名称，只能包含六种字符：英文字母、数字、中文、中划线、下划线和点号。作业名称不能重复。
	Name string `json:"name"`

	// 节点清单
	Nodes []Node `json:"nodes"`

	Schedule *Schedule `json:"schedule"`

	// 作业参数清单
	Params *[]JobParam `json:"params,omitempty"`

	// 日志路径
	LogPath *string `json:"log_path,omitempty"`

	// 目录路径
	Directory *string `json:"directory,omitempty"`

	// 作业类型:  - REAL_TIME: 实时处理  - BATCH: 批处理
	ProcessType JobInfoRequestProcessType `json:"process_type"`

	// 是否选择单任务，默认为false
	SingleNodeJobFlag *bool `json:"single_node_job_flag,omitempty"`

	// 单任务类型
	SingleNodeJobType *JobInfoRequestSingleNodeJobType `json:"single_node_job_type,omitempty"`

	// 创建用户
	CreateUser *string `json:"create_user,omitempty"`

	// 责任人
	Owner *string `json:"owner,omitempty"`

	// 优先级
	Priority *string `json:"priority,omitempty"`

	// 作业最后修改人
	LastUpdateUser *string `json:"last_update_user,omitempty"`

	// 在开启审批开关后，需要填写该字段。表示创建作业的目标状态。  - SAVED: 保存态，表示作业仅保存，无法调度运行，需要提交并审核通过后才能运行。  - SUBMITTED: 提交态，表示作业保存后会自动提交，需要审核通过才能运行。  - PRODUCTION: 生产态，表示作业跳过审批环节，创建后可以直接运行。注意：只有工作空间的管理员用户才能创建生产态的作业。
	TargetStatus *JobInfoRequestTargetStatus `json:"target_status,omitempty"`

	// 作业审批人
	Approvers *[]Approver `json:"approvers,omitempty"`

	BasicConfig *BasicInfo `json:"basic_config,omitempty"`
}

func (o JobInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobInfoRequest struct{}"
	}

	return strings.Join([]string{"JobInfoRequest", string(data)}, " ")
}

type JobInfoRequestProcessType struct {
	value string
}

type JobInfoRequestProcessTypeEnum struct {
	BATCH     JobInfoRequestProcessType
	REAL_TIME JobInfoRequestProcessType
}

func GetJobInfoRequestProcessTypeEnum() JobInfoRequestProcessTypeEnum {
	return JobInfoRequestProcessTypeEnum{
		BATCH: JobInfoRequestProcessType{
			value: "BATCH",
		},
		REAL_TIME: JobInfoRequestProcessType{
			value: "REAL_TIME",
		},
	}
}

func (c JobInfoRequestProcessType) Value() string {
	return c.value
}

func (c JobInfoRequestProcessType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobInfoRequestProcessType) UnmarshalJSON(b []byte) error {
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

type JobInfoRequestSingleNodeJobType struct {
	value string
}

type JobInfoRequestSingleNodeJobTypeEnum struct {
	DLI_SQL        JobInfoRequestSingleNodeJobType
	DWS_SQL        JobInfoRequestSingleNodeJobType
	HIVE_SQL       JobInfoRequestSingleNodeJobType
	SPARK_SQL      JobInfoRequestSingleNodeJobType
	RDS_SQL        JobInfoRequestSingleNodeJobType
	DORIS_SQL      JobInfoRequestSingleNodeJobType
	ASSIGNMENT     JobInfoRequestSingleNodeJobType
	BRANCH         JobInfoRequestSingleNodeJobType
	MERGE          JobInfoRequestSingleNodeJobType
	DATA_MIGRATION JobInfoRequestSingleNodeJobType
	MRS_FLINK      JobInfoRequestSingleNodeJobType
	FLINK_SQL      JobInfoRequestSingleNodeJobType
	FLINK_JAR      JobInfoRequestSingleNodeJobType
	DLI_SPARK      JobInfoRequestSingleNodeJobType
}

func GetJobInfoRequestSingleNodeJobTypeEnum() JobInfoRequestSingleNodeJobTypeEnum {
	return JobInfoRequestSingleNodeJobTypeEnum{
		DLI_SQL: JobInfoRequestSingleNodeJobType{
			value: "DliSQL",
		},
		DWS_SQL: JobInfoRequestSingleNodeJobType{
			value: "DwsSQL",
		},
		HIVE_SQL: JobInfoRequestSingleNodeJobType{
			value: "HiveSQL",
		},
		SPARK_SQL: JobInfoRequestSingleNodeJobType{
			value: "SparkSQL",
		},
		RDS_SQL: JobInfoRequestSingleNodeJobType{
			value: "RdsSQL",
		},
		DORIS_SQL: JobInfoRequestSingleNodeJobType{
			value: "DorisSQL",
		},
		ASSIGNMENT: JobInfoRequestSingleNodeJobType{
			value: "ASSIGNMENT",
		},
		BRANCH: JobInfoRequestSingleNodeJobType{
			value: "BRANCH",
		},
		MERGE: JobInfoRequestSingleNodeJobType{
			value: "MERGE",
		},
		DATA_MIGRATION: JobInfoRequestSingleNodeJobType{
			value: "DataMigration",
		},
		MRS_FLINK: JobInfoRequestSingleNodeJobType{
			value: "MrsFlink",
		},
		FLINK_SQL: JobInfoRequestSingleNodeJobType{
			value: "FlinkSQL",
		},
		FLINK_JAR: JobInfoRequestSingleNodeJobType{
			value: "FlinkJar",
		},
		DLI_SPARK: JobInfoRequestSingleNodeJobType{
			value: "DLISpark",
		},
	}
}

func (c JobInfoRequestSingleNodeJobType) Value() string {
	return c.value
}

func (c JobInfoRequestSingleNodeJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobInfoRequestSingleNodeJobType) UnmarshalJSON(b []byte) error {
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

type JobInfoRequestTargetStatus struct {
	value string
}

type JobInfoRequestTargetStatusEnum struct {
	SAVED      JobInfoRequestTargetStatus
	SUBMITTED  JobInfoRequestTargetStatus
	PRODUCTION JobInfoRequestTargetStatus
}

func GetJobInfoRequestTargetStatusEnum() JobInfoRequestTargetStatusEnum {
	return JobInfoRequestTargetStatusEnum{
		SAVED: JobInfoRequestTargetStatus{
			value: "SAVED",
		},
		SUBMITTED: JobInfoRequestTargetStatus{
			value: "SUBMITTED",
		},
		PRODUCTION: JobInfoRequestTargetStatus{
			value: "PRODUCTION",
		},
	}
}

func (c JobInfoRequestTargetStatus) Value() string {
	return c.value
}

func (c JobInfoRequestTargetStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobInfoRequestTargetStatus) UnmarshalJSON(b []byte) error {
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
