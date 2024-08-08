package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type JobInfo struct {

	// 作业名称
	Name string `json:"name"`

	// 节点定义
	Nodes []Node `json:"nodes"`

	Schedule *Schedule `json:"schedule"`

	// 作业参数定义
	Params *[]JobParam `json:"params,omitempty"`

	// 作业在目录树上的路径。创建作业时如果路径目录不存在，会自动创建目录，如/dir/a/，默认在根目录/。
	Directory *string `json:"directory,omitempty"`

	// 设置作业的最大超时时间。
	CleanOverdueDays *int32 `json:"cleanOverdueDays,omitempty"`

	// 清除等待的作业。
	CleanWaitingJob *string `json:"cleanWaitingJob,omitempty"`

	// 取值为0和1，1表示空跑，0表示：取消空跑，不设置该参数时，默认为0。
	EmptyRunningJob *string `json:"emptyRunningJob,omitempty"`

	// 作业类型，REAL_TIME： 实时处理，BATCH：批处理
	ProcessType JobInfoProcessType `json:"processType"`

	// 是否选择单任务，默认为false
	SingleNodeJobFlag *bool `json:"singleNodeJobFlag,omitempty"`

	// 单任务类型
	SingleNodeJobType *JobInfoSingleNodeJobType `json:"singleNodeJobType,omitempty"`

	// 作业最后修改人
	LastUpdateUser *string `json:"lastUpdateUser,omitempty"`

	// 作业运行日志存放的OBS路径。
	LogPath *string `json:"logPath,omitempty"`

	BasicConfig *BasicConfig `json:"basicConfig,omitempty"`

	// 在开启审批开关后，需要填写该字段。表示创建作业的目标状态，有三种状态：SAVED、SUBMITTED和PRODUCTION，分别表示作业创建后是保存态，提交态，生产态。
	TargetStatus *JobInfoTargetStatus `json:"targetStatus,omitempty"`

	// 在开启审批开关后，需要填写该字段，表示作业（或脚本）审批人。
	Approvers *[]JobApprover `json:"approvers,omitempty"`
}

func (o JobInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobInfo struct{}"
	}

	return strings.Join([]string{"JobInfo", string(data)}, " ")
}

type JobInfoProcessType struct {
	value string
}

type JobInfoProcessTypeEnum struct {
	BATCH     JobInfoProcessType
	REAL_TIME JobInfoProcessType
}

func GetJobInfoProcessTypeEnum() JobInfoProcessTypeEnum {
	return JobInfoProcessTypeEnum{
		BATCH: JobInfoProcessType{
			value: "BATCH",
		},
		REAL_TIME: JobInfoProcessType{
			value: "REAL_TIME",
		},
	}
}

func (c JobInfoProcessType) Value() string {
	return c.value
}

func (c JobInfoProcessType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobInfoProcessType) UnmarshalJSON(b []byte) error {
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

type JobInfoSingleNodeJobType struct {
	value string
}

type JobInfoSingleNodeJobTypeEnum struct {
	DLI_SQL    JobInfoSingleNodeJobType
	DWS_SQL    JobInfoSingleNodeJobType
	HIVE_SQL   JobInfoSingleNodeJobType
	SPARK_SQL  JobInfoSingleNodeJobType
	RDS_SQL    JobInfoSingleNodeJobType
	DORIS_SQL  JobInfoSingleNodeJobType
	ASSIGNMENT JobInfoSingleNodeJobType
	BRANCH     JobInfoSingleNodeJobType
	MERGE      JobInfoSingleNodeJobType
	NORMAL_JOB JobInfoSingleNodeJobType
	ONE_CLICK  JobInfoSingleNodeJobType
	MRS_FLINK  JobInfoSingleNodeJobType
	FLINK_SQL  JobInfoSingleNodeJobType
	FLINK_JAR  JobInfoSingleNodeJobType
	DLI_SPARK  JobInfoSingleNodeJobType
}

func GetJobInfoSingleNodeJobTypeEnum() JobInfoSingleNodeJobTypeEnum {
	return JobInfoSingleNodeJobTypeEnum{
		DLI_SQL: JobInfoSingleNodeJobType{
			value: "DliSQL",
		},
		DWS_SQL: JobInfoSingleNodeJobType{
			value: "DwsSQL",
		},
		HIVE_SQL: JobInfoSingleNodeJobType{
			value: "HiveSQL",
		},
		SPARK_SQL: JobInfoSingleNodeJobType{
			value: "SparkSQL",
		},
		RDS_SQL: JobInfoSingleNodeJobType{
			value: "RdsSQL",
		},
		DORIS_SQL: JobInfoSingleNodeJobType{
			value: "DorisSQL",
		},
		ASSIGNMENT: JobInfoSingleNodeJobType{
			value: "ASSIGNMENT",
		},
		BRANCH: JobInfoSingleNodeJobType{
			value: "BRANCH",
		},
		MERGE: JobInfoSingleNodeJobType{
			value: "MERGE",
		},
		NORMAL_JOB: JobInfoSingleNodeJobType{
			value: "NormalJob",
		},
		ONE_CLICK: JobInfoSingleNodeJobType{
			value: "OneClick",
		},
		MRS_FLINK: JobInfoSingleNodeJobType{
			value: "MrsFlink",
		},
		FLINK_SQL: JobInfoSingleNodeJobType{
			value: "FlinkSQL",
		},
		FLINK_JAR: JobInfoSingleNodeJobType{
			value: "FlinkJar",
		},
		DLI_SPARK: JobInfoSingleNodeJobType{
			value: "DLISpark",
		},
	}
}

func (c JobInfoSingleNodeJobType) Value() string {
	return c.value
}

func (c JobInfoSingleNodeJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobInfoSingleNodeJobType) UnmarshalJSON(b []byte) error {
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

type JobInfoTargetStatus struct {
	value string
}

type JobInfoTargetStatusEnum struct {
	SAVED      JobInfoTargetStatus
	SUBMITTED  JobInfoTargetStatus
	PRODUCTION JobInfoTargetStatus
}

func GetJobInfoTargetStatusEnum() JobInfoTargetStatusEnum {
	return JobInfoTargetStatusEnum{
		SAVED: JobInfoTargetStatus{
			value: "SAVED",
		},
		SUBMITTED: JobInfoTargetStatus{
			value: "SUBMITTED",
		},
		PRODUCTION: JobInfoTargetStatus{
			value: "PRODUCTION",
		},
	}
}

func (c JobInfoTargetStatus) Value() string {
	return c.value
}

func (c JobInfoTargetStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobInfoTargetStatus) UnmarshalJSON(b []byte) error {
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
